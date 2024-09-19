package router

import (
	"vue3-admin-template/internal/admin/controller"

	"github.com/gin-gonic/gin"
)

func initRouterProPrice(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	//authPubRouter := publicRouter.Group("user") //.Use(middleware.OperationRecord())
	authPriRouter := privateGroup

	controller := controller.ControlProPrice{}

	{
		// add
		authPriRouter.POST("/price", controller.Create) // 新建报告的数据结构
		// del
		authPriRouter.DELETE("/price", controller.DeleteMany) // 批量删除报告的数据结构
		// change
		authPriRouter.PATCH("/price/:id", controller.PatchOne) // 更新报告的数据结构
		// get
		authPriRouter.GET("/price", controller.GetPage) // 获取报告的数据结构列表
	}

}
