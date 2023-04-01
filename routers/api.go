package routers

import (
	"mh-go/controllers/api"
	"mh-go/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	routes := r.Group("/api")
	{
		routes.POST("/login", api.LoginController{}.Login)

		user := routes.Group("/user", middlewares.JWTMiddleware())
		{
			user.POST("/info", api.UserController{}.UserInfo)
		}
	}
}
