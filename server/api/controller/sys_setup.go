package controller

import (
	"yc-webreport-server/model"

	"github.com/gin-gonic/gin"
)

type SysSetupControl struct {
	BaseController
}

var modelSetup model.SysSetup

func (control *SysSetupControl) PatchOne(ctx *gin.Context) {
	id := ctx.Param("id")
	// 把id 转为 uint

	req := PatchReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	err = modelSetup.PatchOne(id, req.Field, req.Data)
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	control.RetOK(ctx)
}

func (control *SysSetupControl) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")

	item, err := modelSetup.GetOne(id)
	if err != nil {
		item, err = modelSetup.GetOneByName(id)
		if err != nil {
			control.RetError(ctx, ERROR, err.Error())
			return
		}
	}

	control.RetOkData(ctx, item)
}

func (control *SysSetupControl) GetList(ctx *gin.Context) {

	items, _, err := modelSetup.GetList()
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	control.RetOkList(ctx, items)
}
