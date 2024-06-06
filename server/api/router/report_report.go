package router

import (
	"yc-webreport-server/api/controller"

	"github.com/gin-gonic/gin"
)

func initRouterReportReport(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	authPubRouter := publicRouter.Group("report") //.Use(middleware.OperationRecord())
	authPriRouter := privateGroup.Group("report")

	baseApi := controller.ReportReportControl{}

	authPubRouter.POST("report", baseApi.Create)
	authPriRouter.DELETE("report", baseApi.DeleteMany)
	authPriRouter.DELETE("report/:id", baseApi.DeleteOne)
	authPriRouter.PATCH("report/:id", baseApi.PatchOne)
	authPubRouter.PUT("report/:id", baseApi.PutOne)
	authPriRouter.GET("report/:id", baseApi.GetOne)
	authPriRouter.GET("report", baseApi.GetList)

}
