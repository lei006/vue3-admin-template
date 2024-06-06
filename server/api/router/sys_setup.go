package router

import (
	"yc-webreport-server/api/controller"

	"github.com/gin-gonic/gin"
)

func initRouterSysSetup(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	optionRouter := privateGroup

	controller := controller.SysSetupControl{}

	{
		// add
		//optionRouter.POST("/setup", controller.Create) // 新建报告的数据结构
		// del
		//authPriRouter.DELETE("/user/:id", controller.DeleteOne) // 删除报告的数据结构
		//optionRouter.DELETE("/setup", controller.DeleteMany) // 批量删除报告的数据结构
		// change
		//optionRouter.PUT("/setup/:id", controller.PutOne)     // 更新报告的数据结构
		optionRouter.PATCH("/setup/:id", controller.PatchOne) // 更新报告的数据结构
		// get
		publicRouter.GET("/setup/:id", controller.GetOne) // 根据ID获取报告的数据结构
		publicRouter.GET("/setup", controller.GetList)    // 获取报告的数据结构列表
	}

}
