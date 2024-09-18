package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"vue3-admin-template/internal/config"
	"vue3-admin-template/internal/db_model"

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

	items = append(items, db_model.SysAbout{Title: "已授权", Data: fmt.Sprintf("%t", config.LicenseCheck), Desc: ""})
	if config.LicenseCheck {
		items = append(items, db_model.SysAbout{Title: "限制数量1", Data: fmt.Sprintf("%d", config.LicenseLimit0), Desc: ""})
		items = append(items, db_model.SysAbout{Title: "限制数量2", Data: fmt.Sprintf("%d", config.LicenseLimit1), Desc: ""})
		items = append(items, db_model.SysAbout{Title: "限制数量3", Data: fmt.Sprintf("%d", config.LicenseLimit2), Desc: ""})
	}

	RetPage(ctx, items, total+4)

}
