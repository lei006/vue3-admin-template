package controller

import (
	"yc-webreport-server/api/model"
	"yc-webreport-server/api/utils"

	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
	"go.uber.org/zap"
)

type ReportItemControl struct {
	BaseController
}

var ReportItemController = new(ReportItemControl)

func (control *ReportItemControl) Create(ctx *gin.Context) {
	var report_item model.ReportItem
	err := ctx.ShouldBindJSON(&report_item)
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}
	
	verify := utils.Rules{
		"name":  {utils.NotEmpty()},
		"title": {utils.NotEmpty()},
		"data":  {utils.NotEmpty()},
	}
	if err := utils.Verify(report_item, verify); err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	if err := model.InsReportItem.Create(&report_item); err != nil {
		zlog.Error("创建失败!", zap.Error(err))
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	control.RetOkData(ctx, report_item)
}

// DeleteReportStruct 删除报告的数据结构
func (control *ReportItemControl) DeleteOne(c *gin.Context) {

	id := c.Param("id")

	if err := model.InsReportItem.DeleteOne(id); err != nil {
		zlog.Error("delete one error:", zap.Error(err))
		control.RetError(c, ERROR, err.Error())
		return
	}

	control.RetOK(c)
}

// DeleteReportStructByIds 批量删除报告的数据结构
// @Tags ReportStruct
// @Summary 批量删除报告的数据结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除报告的数据结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /reportStruct/deleteReportStructByIds [delete]
func (control *ReportItemControl) DeleteMany(c *gin.Context) {
	var ids []uint
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		control.RetError(c, ERROR, err.Error())
		return
	}
	if err := model.InsReportItem.DeleteMany(ids); err != nil {
		zlog.Error("批量删除失败!", zap.Error(err))
		zlog.Error("获取失败!", zap.Error(err))

		control.RetError(c, ERROR, err.Error())
		return

	}

	control.RetOK(c)

}

// UpdateReportStruct 更新报告的数据结构
func (control *ReportItemControl) PutOne(ctx *gin.Context) {

	var reportItem model.ReportItem
	err := ctx.ShouldBindJSON(&reportItem)
	if err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}
	verify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	if err := utils.Verify(reportItem, verify); err != nil {
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	if err := model.InsReportItem.UpdateOne(&reportItem); err != nil {
		zlog.Error("更新失败!", zap.Error(err))
		control.RetError(ctx, ERROR, err.Error())
		return

	}
	control.RetOK(ctx)
}

func (control *ReportItemControl) PatchOne(ctx *gin.Context) {
	/*
		var reportStruct model.ReportFieldLayout
		err := ctx.ShouldBindJSON(&reportStruct)
		if err != nil {
			control.RetError(ctx, ERROR, err.Error())
			return
		}
		verify := utils.Rules{
			"ID": {utils.NotEmpty()},
		}
		if err := utils.Verify(reportStruct, verify); err != nil {
			control.RetError(ctx, ERROR, err.Error())
			return
		}

		if err := service.ServerReportFieldLayout.UpdateReportFieldLayout(reportStruct); err != nil {
			zlog.Error("更新失败!", zap.Error(err))
			control.RetError(ctx, ERROR, err.Error())
			return
		}
	*/
	control.RetOK(ctx)
}

// FindReportStruct 用id查询报告的数据结构
func (control *ReportItemControl) GetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := model.InsReportItem.GetOne(id)
	if err != nil {
		zlog.Error("查询失败!", zap.Error(err))
		control.RetError(ctx, ERROR, err.Error())
		return
	}
	control.RetOkData(ctx, data)
}

// GetReportStructList 分页获取报告的数据结构列表

func (control *ReportItemControl) GetList(ctx *gin.Context) {

	data_list, _, err := model.InsReportItem.GetList()
	if err != nil {
		zlog.Error("get report setup list error: ", zap.Error(err))
		control.RetError(ctx, ERROR, err.Error())
		return
	}

	control.ReturnList(ctx, data_list)
}
