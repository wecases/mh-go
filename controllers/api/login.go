package api

import (
	"mh-go/middlewares"
	"mh-go/models"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Login(c *gin.Context) {
	var body struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		con.Error("无效的请求正文", c)
		return
	}

	user, err := models.GetUser(&models.User{Name: body.Name})

	if err != nil {
		con.Error("用户名或密码不正确", c)
		return
	}

	token, err := middlewares.GetToken(user)
	if err != nil {
		con.Error("token生成失败", c)
		return
	}

	con.Success("登录成功", gin.H{
		"token": token,
		"user":  user.Sanitize(),
	}, c)
}
