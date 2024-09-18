package controller

import (
	"fmt"
	"net/http"
	"vue3-admin-template/internal/db_model"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sohaha/zlsgo/zlog"
	"go.uber.org/zap"
)

type SysAuthAdminControl struct {
	ControllerBase
}

func (control *SysAuthAdminControl) Login(ctx *gin.Context) {

	//key := ctx.ClientIP()

	type Login struct {
		Username  string `json:"username"`  // 用户名
		Password  string `json:"password"`  // 密码
		Captcha   string `json:"captcha"`   // 验证码
		CaptchaId string `json:"captchaId"` // 验证码ID
	}

	login := Login{}
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "")
		return
	}

	if CaptchaEnable {
		bret := base64Captcha.DefaultMemStore.Verify(login.CaptchaId, login.Captcha, true)
		if !bret {
			RetErr(ctx, http.StatusBadRequest, "验证码错误")
			return
		}
	}

	modelAdmin := &db_model.SysAdmin{}
	user_info, err := modelAdmin.GetOneByUsername(login.Username)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "用户名或密码错误")
		return
	}

	if user_info.Password != login.Password {
		RetErr(ctx, http.StatusBadRequest, "用户名或密码错误")
		return
	}

	if user_info.IsDisable {
		RetErr(ctx, http.StatusBadRequest, "用户已停用")
		return
	}

	new_token := RandomString(32, true, true, false)

	// 把 user_info.ID 转为字符串
	tmp_id := fmt.Sprintf("%d", user_info.ID)

	err = modelAdmin.PatchOne(tmp_id, "token", new_token)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "请求token出错")
		return
	}

	user_info.Token = new_token
	user_info.Password = ""

	RetData(ctx, user_info)
}

func (control *SysAuthAdminControl) Logout(ctx *gin.Context) {

	token := ctx.Request.Header.Get("token")

	modelAdmin := &db_model.SysAdmin{}
	user_info, err := modelAdmin.GetOneByToken(token)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "未找到用户信息")
		return
	}

	tmp_id := fmt.Sprintf("%d", user_info.ID)
	err = modelAdmin.PatchOne(tmp_id, "token", "")
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "请求出错")
		return
	}

	RetOK(ctx)
}

func (control *SysAuthAdminControl) Info(ctx *gin.Context) {

	token := ctx.Request.Header.Get("token")

	modelAdmin := &db_model.SysAdmin{}
	user_info, err := modelAdmin.GetOneByToken(token)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "未找到token用户信息")
		return
	}

	RetData(ctx, user_info)
}

func (control *SysAuthAdminControl) SetPassword(ctx *gin.Context) {

	token := ctx.Request.Header.Get("token")

	type PasswordResponse struct {
		Old string `json:"old"`
		New string `json:"new"`
	}

	passwordRes := PasswordResponse{}
	err := ctx.ShouldBindJSON(&passwordRes)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	modelUser := &db_model.SysAdmin{}
	user_info, err := modelUser.GetOneByToken(token)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "未找到用户"+err.Error())
		return
	}
	zlog.Debug("passwordRes", passwordRes)
	zlog.Debug("user_info", user_info)
	if user_info.Password != passwordRes.Old {
		RetErr(ctx, http.StatusBadRequest, "原始密码错误")
		return
	}

	err = modelUser.PatchOne(fmt.Sprintf("%d", user_info.ID), "password", passwordRes.New)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "修改密码出错")
		return
	}
}

func (control *SysAuthAdminControl) Regedit(ctx *gin.Context) {

	RetOK(ctx)
}

// 验证码
func (control *SysAuthAdminControl) Captcha(ctx *gin.Context) {

	driver := base64Captcha.NewDriverDigit(CaptchaImgHeight, CaptchaImgWidth, CaptchaKeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		zlog.Error("验证码获取失败!", zap.Error(err))
		RetErr(ctx, http.StatusBadRequest, "验证码获取失败!")
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

	RetData(ctx, CaptchaResponse{
		Id:      id,
		PicPath: b64s,
		Length:  CaptchaKeyLong,
		Height:  CaptchaImgHeight,
		Width:   CaptchaImgWidth,
		Enable:  CaptchaEnable,
	})
}
