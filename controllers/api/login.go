package api

import (
	"mh-go/logic"
	"regexp"

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

	if !regexp.MustCompile(`/^(?:(?:\\+|00)86)?1\\d{11}$/`).MatchString(data.Phone) {
		con.Error("手机号码格式错误", c)
		return
	}

	// 调用注册逻辑
	_, err := logic.Register(data)
	if err != nil {
		con.Error("注册失败", c)
		return
	}

	con.Success("注册成功", nil, c)
}

// 登录
func (con LoginController) Login(c *gin.Context) {
	var data logic.LoginParams

	if err := c.ShouldBind(&data); err != nil {
		con.Error("无效的请求参数", c)
		return
	}

	if !regexp.MustCompile(`/^(?:(?:\\+|00)86)?1\\d{11}$/`).MatchString(data.Phone) {
		con.Error("手机号码格式错误", c)
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

	// user.Sanitize()

	// 返回结果
	con.Success("登录成功", gin.H{
		"token": token,
		"user":  user,
	}, c)
}
