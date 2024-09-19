package controller

import (
	"fmt"
	"net/http"
	"vue3-admin-template/internal/db_model"

	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
)

type ControlProPrice struct {
	ControllerBase
}

func (control *ControlProPrice) Create(ctx *gin.Context) {
	var model_struct db_model.ModelPriceStruct

	err := ctx.ShouldBindJSON(&model_struct)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// 判断项目是否存在
	var model_project_struct db_model.ModelProjectStruct
	_, err = model_project_struct.GetOne(fmt.Sprintf("%d", model_struct.ProjectId))
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, "project not found")
		return
	}

	err = model_struct.Create(&model_struct)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetData(ctx, model_struct)
}

func (control *ControlProPrice) DeleteMany(ctx *gin.Context) {
	var ids []int
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var model_struct db_model.ModelPriceStruct
	err = model_struct.DeleteMany(ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetOK(ctx)
}

func (control *ControlProPrice) PatchOne(ctx *gin.Context) {
	id := ctx.Param("id")
	// 把id 转为 uint

	req := FieldDataRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var model_struct db_model.ModelPriceStruct
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

func (control *ControlProPrice) GetPage(ctx *gin.Context) {

	// 从参数中获取分页信息
	project_id := ctx.Query("project_id")

	var err error

	var model_struct db_model.ModelPriceStruct
	items, total, err := model_struct.GetPage(project_id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetPage(ctx, items, total)

}
