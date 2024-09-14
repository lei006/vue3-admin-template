package router

import (
	"vue3-admin-template/internal/admin/controller"

	"github.com/gin-gonic/gin"
)

func initRouterSysAdmin(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	//authPubRouter := publicRouter.Group("user") //.Use(middleware.OperationRecord())
	authPriRouter := privateGroup

	controller := controller.SysAdminControl{}

	{
		// add
		authPriRouter.POST("/admin", controller.Create) // 新建报告的数据结构
		// del
		authPriRouter.DELETE("/admin/:id", controller.DeleteOne) // 删除报告的数据结构
		authPriRouter.DELETE("/admin", controller.DeleteMany)    // 批量删除报告的数据结构
		// change
		authPriRouter.PUT("/admin/:id", controller.PutOne)     // 更新报告的数据结构
		authPriRouter.PATCH("/admin/:id", controller.PatchOne) // 更新报告的数据结构
		// get
		authPriRouter.GET("/admin/:id", controller.GetOne) // 根据ID获取报告的数据结构
		authPriRouter.GET("/admin", controller.GetPage)    // 获取报告的数据结构列表
	}

}
