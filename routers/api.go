package routers

import (
	"mh-go/controllers/api"
	"mh-go/middlewares"

	"github.com/gin-gonic/gin"
)

func ApiRouters(r *gin.Engine) {
	r.Use(middlewares.Cors())
	routes := r.Group("/api")
	{
		routes.POST("/login", api.LoginController{}.Login)
		routes.POST("/register", api.LoginController{}.Register)

		user := routes.Group("/user", middlewares.JWTMiddleware())
		{
			user.POST("/info", api.UserController{}.UserInfo)
		}
	}
}
