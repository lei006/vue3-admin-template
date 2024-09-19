package router

import (
	"vue3-admin-template/internal/admin/controller"

	"github.com/gin-gonic/gin"
)

func initRouterProProject(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	//authPubRouter := publicRouter.Group("user") //.Use(middleware.OperationRecord())
	authPriRouter := privateGroup

	controller := controller.ControlProProject{}

	{
		// add
		authPriRouter.POST("/project", controller.Create) // 新建报告的数据结构
		// del
		authPriRouter.DELETE("/project", controller.DeleteMany) // 批量删除报告的数据结构
		// change
		authPriRouter.PATCH("/project/:id", controller.PatchOne) // 更新报告的数据结构
		// get
		authPriRouter.GET("/project/:id", controller.GetOne) // 根据ID获取报告的数据结构
		authPriRouter.GET("/project", controller.GetPage)    // 获取报告的数据结构列表
	}

}
