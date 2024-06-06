package router

import (
	"yc-webreport-server/api/controller"

	"github.com/gin-gonic/gin"
)

func initRouterReportPrint(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	authPubRouter := publicRouter.Group("report") //.Use(middleware.OperationRecord())
	authPriRouter := privateGroup.Group("report")

	baseApi := controller.ReportPrintControl{}

	authPubRouter.POST("print", baseApi.Create)
	authPriRouter.DELETE("print", baseApi.DeleteMany)
	authPriRouter.DELETE("print/:id", baseApi.DeleteOne)
	//authPriRouter.PATCH("report/:id", baseApi.PatchOne)
	//authPubRouter.PUT("report/:id", baseApi.PutOne)
	authPubRouter.GET("print/:id", baseApi.GetOne)
	authPubRouter.GET("print", baseApi.GetList)

}
