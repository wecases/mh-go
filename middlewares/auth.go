package middlewares

import (
	"errors"
	"mh-go/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// 秘钥
var Secret = []byte("hahahsuadhasjkdhsdiopwqei")

// 定义 token 中的数据结构
type Claims struct {
	jwt.StandardClaims
	User models.User `json:"user"`
}

// 验证 token
func verifyToken(tokenString string) (*jwt.Token, error) {
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return Secret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// 拦截器
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "令牌不存在",
			})
			c.Abort()
			return
		}

		// 解析并验证 token
		token, err := verifyToken(tokenString)
		if err != nil {
			var e error
			switch err {
			case jwt.ErrSignatureInvalid:
				e = errors.New("无效签名")
			default:
				e = err
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": e.Error(),
			})
			c.Abort()
			return
		}

		// 如果 token 验证通过，则将用户信息注入到 context 中，并继续处理请求
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("user", claims.User)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "令牌无效",
			})
			c.Abort()
			return
		}
	}
}
