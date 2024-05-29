package controller

import (
	"github.com/gin-gonic/gin"
)

type SysUserAuthControl struct {
	BaseController
}

var SysUserAuthControler = new(SysUserAuthControl)

func (control *SysUserAuthControl) Login(ctx *gin.Context) {

	control.RetOK(ctx)
}
func (control *SysUserAuthControl) Logout(ctx *gin.Context) {

	control.RetOK(ctx)
}

func (control *SysUserAuthControl) Regedit(ctx *gin.Context) {

	control.RetOK(ctx)
}

// 验证码
func (control *SysUserAuthControl) Captcha(ctx *gin.Context) {

	control.RetOK(ctx)
}
