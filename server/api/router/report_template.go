package router

import (
	"yc-webreport-server/api/controller"

	"github.com/gin-gonic/gin"
)

func initRouterReportTemplate(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	authPubRouter := publicRouter.Group("report") //.Use(middleware.OperationRecord())
	authPriRouter := privateGroup.Group("report")

	baseApi := controller.ReportTemplateControl{}

	authPubRouter.POST("template", baseApi.Create)
	authPriRouter.DELETE("template", baseApi.DeleteMany)
	authPriRouter.DELETE("template/:id", baseApi.DeleteOne)
	authPriRouter.PATCH("template/:id", baseApi.PatchOne)
	authPubRouter.PUT("template/:id", baseApi.PutOne)
	authPriRouter.GET("template/:id", baseApi.GetOne)
	authPubRouter.GET("template", baseApi.GetList)

}
