package routers

import (
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine, e *engine.Engine) {
	BaseRouters(r)
	AdminRouters(r, e)
	ApiRouters(r)
}
