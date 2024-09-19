package router

import (
	"vue3-admin-template/internal/admin/controller"

	"github.com/gin-gonic/gin"
)

func initRouterProOrder(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	//authPubRouter := publicRouter.Group("user") //.Use(middleware.OperationRecord())
	authPriRouter := privateGroup

	controller := controller.ControlProOrder{}

	{
		// get
		authPriRouter.GET("/order", controller.GetPage) // 获取报告的数据结构列表
	}

}
