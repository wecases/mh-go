package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (con BaseController) Success(message string, data interface{}, c *gin.Context) {
	response := gin.H{
		"code":    http.StatusOK,
		"massage": message,
	}

	if data != nil {
		response["data"] = data
	}

	c.JSON(http.StatusOK, response)
}

func (con BaseController) Error(message string, c *gin.Context) {
	response := gin.H{
		"code":    http.StatusInternalServerError,
		"massage": message,
	}

	c.JSON(http.StatusOK, response)
}
