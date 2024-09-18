package controller

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"vue3-admin-template/internal/config"
	"vue3-admin-template/internal/db_model"
	"vue3-admin-template/internal/license"

	"github.com/gin-gonic/gin"
)

type SysAboutControl struct {
	ControllerBase
}

func (control *SysAboutControl) GetPage(ctx *gin.Context) {

	reqPageInfo := db_model.PageInfo{}
	// 从参数中获取分页信息
	page := ctx.Query("page")
	pageSize := ctx.Query("pageSize")
	keyword := ctx.Query("keyword")

	var err error

	// page 转为 int
	reqPageInfo.Page, err = strconv.Atoi(page)
	if err != nil {
		reqPageInfo.Page = 1
	}
	// pageSize 转为 int
	reqPageInfo.PageSize, err = strconv.Atoi(pageSize)
	if err != nil {
		reqPageInfo.PageSize = 10
	}

	// keyword 转为 string
	reqPageInfo.Keyword = keyword

	var modelDb db_model.SysAbout
	items, total, err := modelDb.GetPage(reqPageInfo)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	items = append(items, db_model.SysAbout{Title: "HardSn", Data: config.HardSn, Desc: ""})

	items = append(items, db_model.SysAbout{Title: "已授权", Data: fmt.Sprintf("%t", config.LicenseCheck), Desc: ""})
	if config.LicenseCheck {
		items = append(items, db_model.SysAbout{Title: "文本1", Data: config.Lic.LicenseText0, Desc: ""})
		items = append(items, db_model.SysAbout{Title: "文本2", Data: config.Lic.LicenseText1, Desc: ""})
		items = append(items, db_model.SysAbout{Title: "文本3", Data: config.Lic.LicenseText2, Desc: ""})
		items = append(items, db_model.SysAbout{Title: "限制数量1", Data: fmt.Sprintf("%d", config.Lic.LicenseLimit0), Desc: ""})
		items = append(items, db_model.SysAbout{Title: "限制数量2", Data: fmt.Sprintf("%d", config.Lic.LicenseLimit1), Desc: ""})
		items = append(items, db_model.SysAbout{Title: "限制数量3", Data: fmt.Sprintf("%d", config.Lic.LicenseLimit2), Desc: ""})
	}

	RetPage(ctx, items, total+4)

}
func (control *SysAboutControl) PatchLicense(ctx *gin.Context) {

	// 接收LIC数据
	data, err := ctx.GetRawData()
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// 解码 url
	tmp_data, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// 解码 base64
	decodedStr, err := url.QueryUnescape(string(tmp_data))
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// 验证 LIC
	bret, err := license.VerifyAndUpdate(decodedStr)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if !bret {
		RetErr(ctx, http.StatusBadRequest, "授权失败")
		return
	}

	//保存到lic文件
	err = license.SaveLicenseFile(decodedStr)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetOK(ctx)

}
