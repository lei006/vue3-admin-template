package controller

import (
	"fmt"
	"net/http"
	"time"
	"vue3-admin-template/internal/admin/model"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sohaha/zlsgo/zlog"
	"go.uber.org/zap"
	"golang.org/x/exp/rand"
)

type SysUserAuthControl struct {
	ControllerBase
}

var CaptchaEnable = false
var CaptchaImgWidth = 100
var CaptchaImgHeight = 100
var CaptchaKeyLong = 4

var ControlerUserAuth = new(SysUserAuthControl)

func (control *SysUserAuthControl) Login(ctx *gin.Context) {

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

	modelUser := &model.SysUser{}
	user_info, err := modelUser.GetOneByUsername(login.Username)
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

	err = modelUser.PatchOne(tmp_id, "token", new_token)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "请求token出错")
		return
	}

	user_info.Token = new_token
	user_info.Password = ""

	RetData(ctx, user_info)
}

func (control *SysUserAuthControl) Logout(ctx *gin.Context) {

	token := ctx.Request.Header.Get("x-token")

	modelUser := &model.SysUser{}
	user_info, err := modelUser.GetOneByToken(token)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "未找到用户信息")
		return
	}

	tmp_id := fmt.Sprintf("%d", user_info.ID)
	err = modelUser.PatchOne(tmp_id, "token", "")
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "请求出错")
		return
	}

	RetOK(ctx)
}

func (control *SysUserAuthControl) Info(ctx *gin.Context) {

	token := ctx.Request.Header.Get("x-token")

	modelUser := &model.SysUser{}
	user_info, err := modelUser.GetOneByToken(token)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "未找到token用户信息")
		return
	}

	RetData(ctx, user_info)
}

func (control *SysUserAuthControl) SetPassword(ctx *gin.Context) {

	type PasswordResponse struct {
		Old string `json:"old_password"`
		New string `json:"new_password"`
	}

	passwordRes := PasswordResponse{}
	err := ctx.ShouldBindJSON(&passwordRes)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := GetUserID(ctx)
	id_str := fmt.Sprintf("%d", id)

	user_info, err := modelUser.GetOne(id_str)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "未找到用户"+err.Error())
		return
	}

	if user_info.Password != passwordRes.Old {
		RetErr(ctx, http.StatusBadRequest, "原始密码错误")
		return
	}

	err = modelUser.PatchOne(id_str, "password", passwordRes.New)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "修改密码出错")
		return
	}

	RetOK(ctx)
}

func (control *SysUserAuthControl) Regedit(ctx *gin.Context) {

	RetOK(ctx)
}

// 验证码
func (control *SysUserAuthControl) Captcha(ctx *gin.Context) {

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

// RandomString 生成指定长度和字符集的随机字符串
func RandomString(length int, numbers, letters, specials bool) string {
	rand.Seed(uint64(time.Now().UnixNano()))

	var charSet []rune

	if numbers {
		charSet = append(charSet, '0', '1', '2', '3', '4', '5', '6', '7', '8', '9')
	}
	if letters {
		charSet = append(charSet, 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z')
	}
	if specials {
		charSet = append(charSet, '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ';', ':', ',', '.', '<', '>', '/', '?', '`', '~')
	}

	if len(charSet) == 0 {
		panic("At least one character type (numbers, letters, or specials) must be enabled.")
	}

	result := make([]rune, length)
	for i := range result {
		result[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(result)
}
