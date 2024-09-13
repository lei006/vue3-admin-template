package router

import (
	"vue3-admin-template/internal/admin/controller"

	"github.com/gin-gonic/gin"
)

func initRouterSysOption(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {
	authPriRouter := privateGroup

	controller := controller.SysRecordControl{}

	{
		authPriRouter.DELETE("/option", controller.DeleteMany) // 批量删除报告的数据结构

		authPriRouter.GET("/option", controller.GetPage) // 获取报告的数据结构列表
	}

}
