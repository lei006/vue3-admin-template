package controller

import (
	"fmt"
	"yc-webreport-server/api/utils"
	"yc-webreport-server/config"
	"yc-webreport-server/model"

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

	modelUser := &model.SysUser{}
	user_info, err := modelUser.GetOneByUsername(login.Username)
	if err != nil {
		zlog.Debug("GetOneByUsername:", err.Error())
		control.RetErrorMessage(ctx, "用户名或密码错误")
		return
	}

	if user_info.Password != login.Password {
		zlog.Debug("password error:", user_info.Password, login.Password)
		control.RetErrorMessage(ctx, "用户名或密码错误")
		return
	}

	if user_info.Disenable {
		zlog.Debug("user disenable error:", user_info.Password, login.Password)
		control.RetErrorMessage(ctx, "用户已禁用")
		return
	}

	new_token := utils.RandomString(32, true, true, false)

	// 把 user_info.ID 转为字符串
	tmp_id := fmt.Sprintf("%d", user_info.ID)

	err = modelUser.PatchOne(tmp_id, "token", new_token)
	if err != nil {
		zlog.Debug("set token error:", zap.Error(err))
		control.RetErrorMessage(ctx, "请求出错")
		return
	}

	user_info.Token = new_token
	user_info.Password = ""

	control.RetOkData(ctx, user_info)
}

func (control *SysUserAuthControl) Logout(ctx *gin.Context) {

	token := ctx.Request.Header.Get("x-token")

	modelUser := &model.SysUser{}
	user_info, err := modelUser.GetOneByToken(token)
	if err != nil {
		zlog.Debug("未找到用户:", err.Error())
		control.RetErrorMessage(ctx, "未找到用户信息")
		return
	}

	tmp_id := fmt.Sprintf("%d", user_info.ID)
	err = modelUser.PatchOne(tmp_id, "token", "")
	if err != nil {
		zlog.Debug("set token error:", zap.Error(err))
		control.RetErrorMessage(ctx, "请求出错")
		return
	}

	control.RetOK(ctx)
}

func (control *SysUserAuthControl) Info(ctx *gin.Context) {

	/*
		type TokenReq struct {
			Token string `json:"token"` // 用户名
		}
		req_token := TokenReq{}
		err := ctx.ShouldBindJSON(&req_token)
		if err != nil {
			zlog.Debug("token:", zap.Error(err))
			control.RetErrorParam(ctx, "")
			return
		}
	*/
	token := ctx.Request.Header.Get("x-token")

	modelUser := &model.SysUser{}
	user_info, err := modelUser.GetOneByToken(token)
	if err != nil {
		zlog.Debug("未找到用户:", err.Error())
		control.RetErrorMessage(ctx, "未找到用户信息")
		return
	}

	control.RetOkData(ctx, user_info)
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
