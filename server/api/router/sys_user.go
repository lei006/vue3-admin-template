package router

import (
	"yc-webreport-server/api/controller"

	"github.com/gin-gonic/gin"
)

func initRouterSysUser(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	//authPubRouter := publicRouter.Group("user") //.Use(middleware.OperationRecord())
	authPriRouter := privateGroup

	controller := controller.SysUserControl{}

	{
		// add
		authPriRouter.POST("/user", controller.Create) // 新建报告的数据结构
		// del
		//authPriRouter.DELETE("/user/:id", controller.DeleteOne) // 删除报告的数据结构
		authPriRouter.DELETE("/user", controller.DeleteMany) // 批量删除报告的数据结构
		// change
		authPriRouter.PUT("/user/:id", controller.PutOne)     // 更新报告的数据结构
		authPriRouter.PATCH("/user/:id", controller.PatchOne) // 更新报告的数据结构
		// get
		authPriRouter.GET("/user/:id", controller.GetOne) // 根据ID获取报告的数据结构
		authPriRouter.GET("/user", controller.GetList)    // 获取报告的数据结构列表
	}

	/*
		userRouter := privateGroup.Group("user") //.Use(middleware.OperationRecord())
		userRouterWithoutRecord := privateGroup.Group("user")

		baseApi := controller.BaseApiApp

		{
			userRouter.POST("admin_register", baseApi.Register)               // 管理员注册账号
			userRouter.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
			userRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
			userRouter.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
			userRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
			userRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
			userRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
			userRouter.POST("resetPassword", baseApi.ResetPassword)           // 设置用户权限组
		}
		{
			userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
			userRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)  // 获取自身信息
		}
	*/
}
