package controller

import (
	"net/http"
	"strconv"
	"vue3-admin-template/internal/admin/model"

	"github.com/gin-gonic/gin"
)

type SysRecordControl struct {
	ControllerBase
}

var SysRecordControler = new(SysRecordControl)

func (control *SysRecordControl) Create(ctx *gin.Context) {

	RetOK(ctx)
}

func (control *SysRecordControl) DeleteMany(ctx *gin.Context) {
	var ids []int
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var modelOption model.SysOption
	err = modelOption.DeleteMany(ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetOK(ctx)

}

func (control *SysRecordControl) GetPage(ctx *gin.Context) {

	reqPageInfo := model.PageInfo{}
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

	var modelOption model.SysOption
	items, total, err := modelOption.GetPage(reqPageInfo)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetPage(ctx, items, total)
}
