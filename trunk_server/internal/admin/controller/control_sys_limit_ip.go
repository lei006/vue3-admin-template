package controller

import (
	"net/http"
	"strconv"
	"vue3-admin-template/internal/admin/model"

	"github.com/gin-gonic/gin"
	"github.com/lei006/zlog"
)

type SysLimitIpControl struct {
	ControllerBase
}

func (control *SysLimitIpControl) Create(ctx *gin.Context) {

	var modelSysLimitIp model.SysLimitIp

	limitIp := model.SysLimitIp{}
	err := ctx.ShouldBindJSON(&limitIp)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	val, err := modelSysLimitIp.Create(&limitIp)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	zlog.Info("增加限制IP成功", val)

	RetData(ctx, val)
}

func (control *SysLimitIpControl) DeleteMany(ctx *gin.Context) {
	var ids []int
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var modelLimitIp model.SysLimitIp
	err = modelLimitIp.DeleteMany(ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetOK(ctx)

}

func (control *SysLimitIpControl) PatchOne(ctx *gin.Context) {
	id := ctx.Param("id")
	// 把id 转为 uint

	req := FieldDataRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var modelLimitIp model.SysLimitIp

	err = modelLimitIp.PatchOne(id, req.Field, req.Data)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	data, err := modelLimitIp.GetOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	zlog.Debug("patch:", id, req, data)

	RetData(ctx, data)
}

func (control *SysLimitIpControl) GetPage(ctx *gin.Context) {

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

	var modelLimitIp model.SysLimitIp
	items, total, err := modelLimitIp.GetPage(reqPageInfo)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetPage(ctx, items, total)
}
