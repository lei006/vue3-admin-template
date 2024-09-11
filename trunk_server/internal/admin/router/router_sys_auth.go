package router

import (
	"vue3-admin-template/internal/admin/controller"

	"github.com/gin-gonic/gin"
)

func initRouterSysAuthUser(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	authPubRouter := publicRouter.Group("auth/user") //.Use(middleware.OperationRecord())
	authPriRouter := privateGroup.Group("auth/user")

	baseApi := controller.SysAuthUserControl{}

	authPubRouter.POST("login", baseApi.Login)
	authPriRouter.DELETE("logout", baseApi.Logout)
	authPriRouter.PATCH("setpassword", baseApi.SetPassword)
	authPubRouter.POST("regedit", baseApi.Regedit)
	authPriRouter.GET("info", baseApi.Info)
	authPubRouter.GET("captcha", baseApi.Captcha)
}

func initRouterSysAuthAdmin(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	authPubRouter := publicRouter.Group("auth/admin") //.Use(middleware.OperationRecord())
	authPriRouter := privateGroup.Group("auth/admin")

	baseApi := controller.SysAuthAdminControl{}

	authPubRouter.POST("login", baseApi.Login)
	authPriRouter.DELETE("logout", baseApi.Logout)
	authPriRouter.PATCH("setpassword", baseApi.SetPassword)
	authPubRouter.POST("regedit", baseApi.Regedit)
	authPriRouter.GET("info", baseApi.Info)
	authPubRouter.GET("captcha", baseApi.Captcha)
}
