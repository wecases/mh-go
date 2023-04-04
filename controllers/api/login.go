package api

import (
	"mh-go/logic"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

// 注册
func (con LoginController) Register(c *gin.Context) {

	// 绑定参数
	var data logic.RegisterParams

	// 校验参数
	if err := c.ShouldBind(&data); err != nil {
		con.Exception(c, "无效的请求参数")
		return
	}

	// 调用注册逻辑
	if _, err := logic.Register(data); err != nil {
		con.Error(c, "注册失败")
		return
	}

	con.Success(c, "注册成功", nil)
}

// 登录
func (con LoginController) Login(c *gin.Context) {
	var data logic.LoginParams

	if err := c.ShouldBind(&data); err != nil {
		con.Exception(c, "无效的请求参数")
		return
	}

	// 调用登录逻辑
	user, err := logic.Login(data)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	// 生成token
	token, err := logic.GetToken(user)
	if err != nil {
		con.Error(c, "token生成失败")
		return
	}

	// user.Sanitize()

	// 返回结果
	con.Success(c, "登录成功", gin.H{
		"token": token,
		"user":  user,
	})
}
