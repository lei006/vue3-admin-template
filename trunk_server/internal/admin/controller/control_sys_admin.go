package controller

import (
	"net/http"
	"vue3-admin-template/internal/db_model"
	"vue3-admin-template/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
)

type SysAdminControl struct {
	ControllerBase
}

func (control *SysAdminControl) Create(ctx *gin.Context) {

	user_info := db_model.SysAdmin{}
	err := ctx.ShouldBindJSON(&user_info)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	zlog.Debugf("user_info: %+v \n", user_info)
	modelAdmin := &db_model.SysAdmin{}
	err = modelAdmin.Create(&user_info)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetData(ctx, user_info)
}

func (control *SysAdminControl) DeleteOne(ctx *gin.Context) {

	id := ctx.Param("id")

	modelAdmin := &db_model.SysAdmin{}
	err := modelAdmin.DeleteOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetOK(ctx)
}

func (control *SysAdminControl) DeleteMany(ctx *gin.Context) {
	var ids []uint
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	modelAdmin := &db_model.SysAdmin{}
	err = modelAdmin.DeleteMany(ids)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetOK(ctx)
}

func (control *SysAdminControl) PutOne(ctx *gin.Context) {

	var reportItem db_model.SysAdmin
	err := ctx.ShouldBindJSON(&reportItem)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	if err := utils.Verify(reportItem, verify); err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	modelAdmin := &db_model.SysAdmin{}
	if err := modelAdmin.UpdateOne(&reportItem); err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	RetOK(ctx)
}

func (control *SysAdminControl) PatchOne(ctx *gin.Context) {
	id := ctx.Param("id")
	// 把id 转为 uint

	req := FieldDataRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	modelAdmin := &db_model.SysAdmin{}
	err = modelAdmin.PatchOne(id, req.Field, req.Data)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	data, err := modelAdmin.GetOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	zlog.Debug("patch:", id, req, data)

	RetData(ctx, data)
}

func (control *SysAdminControl) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")

	modelAdmin := &db_model.SysAdmin{}
	item, err := modelAdmin.GetOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetData(ctx, item)
}

func (control *SysAdminControl) GetPage(ctx *gin.Context) {

	modelAdmin := &db_model.SysAdmin{}
	items, total, err := modelAdmin.GetPage()
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetPage(ctx, items, total)

}
