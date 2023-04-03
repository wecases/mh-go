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
		con.Error("无效的请求参数", c)
		return
	}

	// 调用注册逻辑
	user, err := logic.Register(data)
	if err != nil {
		con.Error("注册失败", c)
		return
	}

	con.Success("注册成功", gin.H{
		"user": user,
	}, c)
}

// 登录
func (con LoginController) Login(c *gin.Context) {
	var data logic.LoginParams

	if err := c.ShouldBind(&data); err != nil {
		con.Error("无效的请求参数", c)
		return
	}

	// 调用登录逻辑
	user, err := logic.Login(data)
	if err != nil {
		con.Error(err.Error(), c)
		return
	}

	// 生成token
	token, err := logic.GetToken(user)
	if err != nil {
		con.Error("token生成失败", c)
		return
	}

	// 返回结果
	con.Success("登录成功", gin.H{
		"token": token,
		"user":  user,
	}, c)
}
