package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

// 成功响应
func (con BaseController) Success(c *gin.Context, message string, data interface{}) {
	response := gin.H{
		"code":    http.StatusOK,
		"massage": message,
	}

	if data != nil {
		response["data"] = data
	}

	c.JSON(http.StatusOK, response)
}

// 失败响应
func (con BaseController) Error(c *gin.Context, message string) {
	response := gin.H{
		"code":    http.StatusBadRequest,
		"massage": message,
	}

	c.JSON(http.StatusOK, response)
}

// 异常响应
func (con BaseController) Exception(c *gin.Context, message string) {
	response := gin.H{
		"code":    http.StatusInternalServerError,
		"massage": message,
	}

	c.JSON(http.StatusInternalServerError, response)
}
