package router

import (
	"yc-webreport-server/api/controller"

	"github.com/gin-gonic/gin"
)

func initReportItemRouter(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {
	//privateGroupRecord := privateGroup.Use(middleware.OperationRecord())
	privateGroupRecord := publicRouter

	{
		// add
		privateGroupRecord.POST("/report_item", controller.ReportItemController.Create) // 新建报告的数据结构
		// del
		privateGroupRecord.DELETE("/report_item/:id", controller.ReportItemController.DeleteOne) // 删除报告的数据结构
		privateGroupRecord.DELETE("/report_item", controller.ReportItemController.DeleteMany)    // 批量删除报告的数据结构
		// change
		privateGroupRecord.PUT("/report_item/:id", controller.ReportItemController.PutOne)     // 更新报告的数据结构
		privateGroupRecord.PATCH("/report_item/:id", controller.ReportItemController.PatchOne) // 更新报告的数据结构
		// get
		publicRouter.GET("/report_item/:id", controller.ReportItemController.GetOne) // 根据ID获取报告的数据结构
		publicRouter.GET("/report_item", controller.ReportItemController.GetList)    // 获取报告的数据结构列表
	}

}
