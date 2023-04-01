package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (con BaseController) Success(message string, c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"massage": message,
	})
}

func (con BaseController) Error(message string, c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    http.StatusInternalServerError,
		"massage": message,
	})
}
