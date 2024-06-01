package router

import (
	"yc-webreport-server/api/controller"

	"github.com/gin-gonic/gin"
)

func initSysUserAuthRouter(publicRouter *gin.RouterGroup, privateGroup *gin.RouterGroup) {

	authPubRouter := publicRouter.Group("auth") //.Use(middleware.OperationRecord())
	authPriRouter := privateGroup.Group("auth")

	baseApi := controller.ControlerUserAuth

	authPubRouter.POST("login", baseApi.Login)
	authPriRouter.DELETE("logout", baseApi.Logout)
	authPriRouter.PATCH("password", baseApi.SetPassword)
	authPubRouter.POST("regedit", baseApi.Regedit)
	authPriRouter.GET("info", baseApi.Info)
	authPubRouter.GET("captcha", baseApi.Captcha)
}
