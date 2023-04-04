package routers

import (
	"mh-go/controllers/api"
	"mh-go/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	routes := r.Group("/api")
	{
		// 中间件
		routes.Use(middlewares.Cors())
		routes.Use(middlewares.RegexValidator())

		// 登录注册
		routes.POST("/login", api.LoginController{}.Login)
		routes.POST("/register", api.LoginController{}.Register)

		user := routes.Group("/user", middlewares.JWTMiddleware())
		{
			user.POST("/info", api.UserController{}.UserInfo)
		}
	}
}
