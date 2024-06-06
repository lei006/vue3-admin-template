package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
)

type ReportReportControl struct {
	BaseController
}

func (control *ReportReportControl) Create(ctx *gin.Context) {

	control.RetOK(ctx)
}

func (control *ReportReportControl) DeleteOne(ctx *gin.Context) {
	id := ctx.Param("id")

	control.RetOkData(ctx, id)
}

func (control *ReportReportControl) DeleteMany(ctx *gin.Context) {

	var ids []uint
	err := ctx.ShouldBindJSON(&ids)
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	zlog.Debug("delete many:", ids)

	control.RetOK(ctx)
}

func (control *ReportReportControl) PutOne(ctx *gin.Context) {
	id := ctx.Param("id")
	zlog.Debug("PutOne:", id)

	control.RetOK(ctx)
}

func (control *ReportReportControl) PatchOne(ctx *gin.Context) {
	id := ctx.Param("id")
	zlog.Debug("PatchOne:", id)

	control.RetOK(ctx)
}

func (control *ReportReportControl) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	zlog.Debug("GetOne:", id)

	control.RetOK(ctx)
}

func (control *ReportReportControl) GetList(ctx *gin.Context) {
	zlog.Debug("GetList")
	control.RetOkPage(ctx, 100, nil)
}
