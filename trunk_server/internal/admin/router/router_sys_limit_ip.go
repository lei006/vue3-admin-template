package router

import (
	"vue3-admin-template/internal/admin/controller"

	"github.com/gin-gonic/gin"
)

func initRouterSysLimitIp(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {
	authPriRouter := privateGroup

	controller := controller.SysLimitIpControl{}

	{
		// add
		authPriRouter.POST("/limit_ip", controller.Create) // 新建报告的数据结构

		// delete
		authPriRouter.DELETE("/limit_ip", controller.DeleteMany) // 批量删除报告的数据结构

		authPriRouter.PATCH("/limit_ip/:id", controller.PatchOne) // 更新报告的数据结构

		// get
		authPriRouter.GET("/limit_ip", controller.GetPage) // 获取报告的数据结构列表
	}

}
