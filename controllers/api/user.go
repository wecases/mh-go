package api

import (
	"mh-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) UserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	userinfo := models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"userinfo": userinfo,
		"user":     user,
	})
}
