package controller

import (
	"io"
	"net/http"
	"strconv"
	"vue3-admin-template/internal/admin/model"

	"github.com/gin-gonic/gin"
)

type SysSetupControl struct {
	ControllerBase
}

var modelSetup model.SysSetup

func (control *SysSetupControl) PatchOne(ctx *gin.Context) {
	id := ctx.Param("id")
	// 把id 转为 uint

	req := FieldDataRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = modelSetup.PatchOne(id, req.Field, req.Data)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	item, err := modelSetup.GetOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetData(ctx, item)
}

func (control *SysSetupControl) GetOneById(ctx *gin.Context) {
	id := ctx.Param("id")

	item, err := modelSetup.GetOne(id)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetData(ctx, item)
}

func (control *SysSetupControl) GetOneByName(ctx *gin.Context) {
	name := ctx.Param("name")

	item, err := modelSetup.GetOneByName(name)
	if err != nil {
		new_val := &model.SysSetup{Name: name}
		err := modelSetup.Create(new_val)
		if err != nil {
			RetErr(ctx, http.StatusBadRequest, err.Error())
			return
		}

		RetData(ctx, new_val)

		return
	}

	RetData(ctx, item)

}

func (control *SysSetupControl) SetOneByName(ctx *gin.Context) {
	name := ctx.Param("name")
	item, err := modelSetup.GetOneByName(name)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}
	modelSetup.PatchOne(strconv.FormatUint(uint64(item.ID), 10), "data", string(bodyBytes))

	RetData(ctx, item)

}

func (control *SysSetupControl) GetList(ctx *gin.Context) {

	items, _, err := modelSetup.GetList()
	if err != nil {
		RetErr(ctx, http.StatusBadRequest, err.Error())
		return
	}

	RetList(ctx, items)
}
