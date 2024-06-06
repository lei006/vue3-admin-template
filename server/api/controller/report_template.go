package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
)

type ReportTemplateControl struct {
	BaseController
}

func (control *ReportTemplateControl) Create(ctx *gin.Context) {

	control.RetOK(ctx)
}

func (control *ReportTemplateControl) DeleteOne(ctx *gin.Context) {
	id := ctx.Param("id")

	control.RetOkData(ctx, id)
}

func (control *ReportTemplateControl) DeleteMany(ctx *gin.Context) {

	var ids []uint
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	zlog.Debug("delete many:", ids)

	control.RetOK(ctx)
}

func (control *ReportTemplateControl) PutOne(ctx *gin.Context) {
	id := ctx.Param("id")
	zlog.Debug("PutOne:", id)

	control.RetOK(ctx)
}

func (control *ReportTemplateControl) PatchOne(ctx *gin.Context) {
	id := ctx.Param("id")
	zlog.Debug("PatchOne:", id)

	control.RetOK(ctx)
}

func (control *ReportTemplateControl) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	zlog.Debug("GetOne:", id)

	control.RetOK(ctx)
}

func (control *ReportTemplateControl) GetList(ctx *gin.Context) {
	zlog.Debug("GetList")
	control.RetOkList(ctx, nil)
}
