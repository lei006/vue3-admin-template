package router

import (
	"vue3-admin-template/internal/admin/controller"

	"github.com/gin-gonic/gin"
)

func initRouterSysLicense(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	//authPubRouter := publicRouter.Group("user") //.Use(middleware.OperationRecord())
	authPriRouter := privateGroup

	controller := controller.SysLicense{}

	{
		// add
		authPriRouter.POST("/license", controller.Create) // 新建报告的数据结构
		// del
		authPriRouter.DELETE("/license/:id", controller.DeleteOne) // 删除报告的数据结构
		authPriRouter.DELETE("/license", controller.DeleteMany)    // 批量删除报告的数据结构
		// change
		authPriRouter.PATCH("/license/:id", controller.PatchOne) // 更新报告的数据结构
		// get
		authPriRouter.GET("/license/:id", controller.GetOne) // 根据ID获取报告的数据结构
		authPriRouter.GET("/license", controller.GetPage)    // 获取报告的数据结构列表
	}

}
