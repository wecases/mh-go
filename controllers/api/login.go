package api

import (
	"mh-go/middlewares"
	"mh-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Login(c *gin.Context) {
	user, err := models.GetUser()
	if err != nil {
		con.Error("用户不正确", c)
	}

	token, err := middlewares.GetToken(user)

	if err != nil {
		con.Error("token生成失败", c)
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}
