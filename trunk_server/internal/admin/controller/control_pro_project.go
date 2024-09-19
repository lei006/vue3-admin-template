package controller

import (
	"net/http"
	"strconv"
	"vue3-admin-template/internal/db_model"

	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
)

type ControlProProject struct {
	ControllerBase
}

func (control *ControlProProject) Create(ctx *gin.Context) {

	var model_struct db_model.ModelProjectStruct
	err := ctx.ShouldBindJSON(&model_struct)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = model_struct.Create(&model_struct)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetData(ctx, model_struct)
}

func (control *ControlProProject) DeleteMany(ctx *gin.Context) {
	var ids []int
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var model_struct db_model.ModelProjectStruct
	err = model_struct.DeleteMany(ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetOK(ctx)
}

func (control *ControlProProject) PatchOne(ctx *gin.Context) {
	id := ctx.Param("id")
	// 把id 转为 uint

	req := FieldDataRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var model_struct db_model.ModelProjectStruct

	err = model_struct.PatchOne(id, req.Field, req.Data)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	data, err := model_struct.GetOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	zlog.Debug("patch:", id, req, data)

	RetData(ctx, data)
}

func (control *ControlProProject) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")

	var model_struct db_model.ModelProjectStruct

	item, err := model_struct.GetOne(id)
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

func (control *ControlProProject) GetPage(ctx *gin.Context) {

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

	var modelDb db_model.ModelProjectStruct
	items, total, err := modelDb.GetPage(reqPageInfo)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetPage(ctx, items, total)

}
