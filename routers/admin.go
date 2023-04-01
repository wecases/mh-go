package routers

import (
	"mh-go/pages"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.Engine, eng *engine.Engine) {

	eng.HTML("GET", "/admin", pages.GetDashBoard)

	eng.HTMLFile("GET", "/admin/hello", "hello.tmpl", map[string]interface{}{
		"msg": "Hello world",
	})
}
