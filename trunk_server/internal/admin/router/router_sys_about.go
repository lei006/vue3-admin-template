package router

import (
	"vue3-admin-template/internal/admin/controller"

	"github.com/gin-gonic/gin"
)

func initRouterSysAbout(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	//optionRouter := privateGroup

	controller := controller.SysAboutControl{}

	{
		publicRouter.GET("/about", controller.GetPage) // 获取报告的数据结构列表
	}

}
