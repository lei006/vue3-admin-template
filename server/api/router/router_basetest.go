package router

import (
	"yc-webreport-server/api/controller"

	"github.com/gin-gonic/gin"
)

func initRouterBaseTest(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	authPubRouter := publicRouter.Group("base") //.Use(middleware.OperationRecord())
	authPriRouter := privateGroup.Group("base")

	baseApi := controller.ReportDemoControl{}

	authPubRouter.POST("demo", baseApi.Create)
	authPriRouter.DELETE("demo", baseApi.DeleteMany)
	authPriRouter.DELETE("demo/:id", baseApi.DeleteOne)
	authPriRouter.PATCH("demo/:id", baseApi.PatchOne)
	authPubRouter.PUT("demo/:id", baseApi.PutOne)
	authPriRouter.GET("demo/:id", baseApi.GetOne)
	authPubRouter.GET("demo", baseApi.GetList)

}
