package logic

import (
	"errors"
	"mh-go/middlewares"
	"mh-go/models"
	"mh-go/utils"
	"time"

	"github.com/golang-jwt/jwt"
)

// 注册参数
type RegisterParams struct {
	LoginParams
	InviteCode string `form:"invite_code"`
}

// 注册逻辑
func Register(data RegisterParams) (*models.User, error) {
	// 声明用户模型
	user := &models.User{
		Name:     data.Phone,
		Phone:    data.Phone,
		Password: data.Password,
	}

	// 加密密码
	err := user.HashPassword()
	if err != nil {
		return nil, err
	}

	// 判断邀请码是否为空
	if data.InviteCode != "" {
		var inviteUser models.User
		// 根据邀请码获取上级用户
		err := models.DB.Where("invite_code = ?", data.InviteCode).First(&inviteUser).Error
		if err != nil {
			return nil, errors.New("邀请码无效")
		}

		// 设置上级ID
		user.ParentID = inviteUser.ID
		// 设置团队路径
		user.Path = append(inviteUser.Path, inviteUser.ID)
	}

	// 设置用户名为手机号的脱敏形式
	user.Name = user.Phone[:3] + "****" + user.Phone[7:]

	// 生成随机邀请码
	user.InviteCode = utils.GenerateCode(8)

	if err := models.DB.Create(&user); err != nil {
		return nil, err.Error
	}

	return user, nil
}

// 登录参数
type LoginParams struct {
	Phone    string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// 登录逻辑
func Login(data LoginParams) (*models.User, error) {
	var user models.User

	// 根据手机号查询用户
	err := models.DB.Where("phone = ?", data.Phone).First(&user).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 校验密码
	if user.VerifyPassword(data.Password) != nil {
		return nil, errors.New("用户名或密码不正确")
	}

	return &user, nil
}

// 生成 token
func GetToken(user *models.User) (string, error) {
	claims := &middlewares.Claims{
		User: *user,
		StandardClaims: jwt.StandardClaims{
			// 有效期一小时
			NotBefore: time.Now().Unix() - 60,
			// 有效期两小时
			ExpiresAt: time.Now().Unix() + 60*60*2,
			// 发行人
			Issuer: "waset",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(middlewares.Secret)
}
