package router

import (
	"yc-webreport-server/api/controller"

	"github.com/gin-gonic/gin"
)

func initRouterReportPrefield(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	authPubRouter := publicRouter.Group("report") //.Use(middleware.OperationRecord())
	authPriRouter := privateGroup.Group("report")

	baseApi := controller.ReportPrefieldControl{}

	authPubRouter.POST("prefield", baseApi.Create)
	authPriRouter.DELETE("prefield", baseApi.DeleteMany)
	authPriRouter.DELETE("prefield/:id", baseApi.DeleteOne)
	authPriRouter.PATCH("prefield/:id", baseApi.PatchOne)
	authPubRouter.PUT("prefield/:id", baseApi.PutOne)
	authPriRouter.GET("prefield/:id", baseApi.GetOne)
	authPriRouter.GET("prefield", baseApi.GetList)

}
