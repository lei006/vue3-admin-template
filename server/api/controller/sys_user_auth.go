package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
)

type SysUserAuthControl struct {
	BaseController
}

var ControlerUserAuth = new(SysUserAuthControl)

func (control *SysUserAuthControl) Login(ctx *gin.Context) {

	type Login struct {
		Username  string `json:"username"`  // 用户名
		Password  string `json:"password"`  // 密码
		Captcha   string `json:"captcha"`   // 验证码
		CaptchaId string `json:"captchaId"` // 验证码ID
	}

	login := Login{}
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		control.RetErrorParam(ctx, "")
		return
	}
	zlog.Debug("login:", login)
	/*
		err := ctx.ShouldBindJSON(&l)
		//key := ctx.ClientIP()

		if err != nil {
			//baseApi.FailWithMessage(err.Error(), c)
			return
		}*/

	control.RetOK(ctx)
}

func (control *SysUserAuthControl) Logout(ctx *gin.Context) {

	control.RetOK(ctx)
}

func (control *SysUserAuthControl) Info(ctx *gin.Context) {

	control.RetOK(ctx)
}

func (control *SysUserAuthControl) SetPassword(ctx *gin.Context) {

	control.RetOK(ctx)
}

func (control *SysUserAuthControl) Regedit(ctx *gin.Context) {

	control.RetOK(ctx)
}

// 验证码
func (control *SysUserAuthControl) Captcha(ctx *gin.Context) {

	control.RetOK(ctx)
}
