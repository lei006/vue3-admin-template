package controller

import (
	"net/http"
	"strconv"
	"time"
	"vue3-admin-template/internal/config"
	"vue3-admin-template/internal/db_model"
	"vue3-admin-template/internal/license"

	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
)

type SysLicense struct {
	ControllerBase
}

func (control *SysLicense) Create(ctx *gin.Context) {
	var modelUser db_model.SysLicense

	license_info := db_model.SysLicense{}
	err := ctx.ShouldBindJSON(&license_info)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	///////////////////////////////////////////
	// 生成签名数据
	license_info.LicenseAt = time.Now()
	license_info.PubKey = config.LicensePubKey

	sign, err := license.Sign(&license_info.LicenseStruct, config.LicensePriKey)
	if err != nil {
		RetErr(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	license_info.Sign = sign
	license_info.LicenseData, err = db_model.LicenseGetJson(&license_info.LicenseStruct)
	if err != nil {
		RetErr(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	//验证一下
	ret, err := license.VerifyOnly(license_info.LicenseData)
	if err != nil {
		RetErr(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if !ret {
		RetErr(ctx, http.StatusInternalServerError, "license 验证失败")
		return
	}

	err = modelUser.Create(&license_info)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetData(ctx, license_info)
}

func (control *SysLicense) DeleteOne(ctx *gin.Context) {

	id := ctx.Param("id")
	var modelUser db_model.SysLicense

	err := modelUser.DeleteOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetOK(ctx)
}

func (control *SysLicense) DeleteMany(ctx *gin.Context) {
	var ids []int
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var modelUser db_model.SysLicense

	err = modelUser.DeleteMany(ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetOK(ctx)
}

func (control *SysLicense) PatchOne(ctx *gin.Context) {
	id := ctx.Param("id")
	// 把id 转为 uint

	req := FieldDataRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var modelUser db_model.SysLicense

	err = modelUser.PatchOne(id, req.Field, req.Data)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	data, err := modelUser.GetOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	zlog.Debug("patch:", id, req, data)

	RetData(ctx, data)
}

func (control *SysLicense) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")

	var modelUser db_model.SysLicense

	item, err := modelUser.GetOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	/*
		ctx.Header("Content-Disposition", "attachment; filename=license.txt")
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.String(http.StatusOK, item.LicenseData)
	*/
	RetData(ctx, item)
}

func (control *SysLicense) GetPage(ctx *gin.Context) {

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

	var modelDb db_model.SysLicense
	items, total, err := modelDb.GetPage(reqPageInfo)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetPage(ctx, items, total)

}
