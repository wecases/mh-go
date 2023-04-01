package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BaseRouters(r *gin.Engine) {
	routes := r.Group("/")
	{
		routes.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "hello.tmpl", gin.H{})
		})
	}
}
