package controller

import (
	"yc-webreport-server/config"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sohaha/zlsgo/zlog"
	"go.uber.org/zap"
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
		zlog.Debug("login:", zap.Error(err))
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

	driver := base64Captcha.NewDriverDigit(config.ReportCfg.Captcha.ImgHeight, config.ReportCfg.Captcha.ImgWidth, config.ReportCfg.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		zlog.Error("验证码获取失败!", zap.Error(err))
		control.RetErrorMessage(ctx, "验证码获取失败!")
		return
	}

	type CaptchaResponse struct {
		Id      string `json:"id"`
		PicPath string `json:"captcha"`
		Length  int    `json:"length"`
		Height  int    `json:"height"`
		Width   int    `json:"width"`
		Enable  bool   `json:"enable"`
	}

	control.RetOkData(ctx, CaptchaResponse{
		Id:      id,
		PicPath: b64s,
		Length:  config.ReportCfg.Captcha.KeyLong,
		Height:  config.ReportCfg.Captcha.ImgHeight,
		Width:   config.ReportCfg.Captcha.ImgWidth,
		Enable:  config.ReportCfg.Captcha.Enable,
	})
}
