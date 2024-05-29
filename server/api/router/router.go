package router

import (
	"net/http"
	"yc-webreport-server/api/middleware"
	"yc-webreport-server/config"

	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
)

func LoadRouter(engine *gin.Engine) error {

	engine.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求

	{
		// 应映页面
		staticPath := config.ReportCfg.WorkPath + "\\static"
		engine.Static("/static", staticPath)
		zlog.Info("/static -- >", staticPath)
	}

	{
		viewPath := config.ReportCfg.WorkPath + "\\view"
		engine.StaticFile("/index.html", viewPath+"\\index.html")
		engine.StaticFile("/favicon.ico", viewPath+"\\favicon.ico")
		engine.StaticFile("/", viewPath+"\\index.html")
		engine.StaticFS("/assets", http.Dir(viewPath+"\\assets"))
		zlog.Info("/ -- >", viewPath)

	}

	{

		// 检查健康测试
		engine.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	publicGroup := engine.Group(config.ReportCfg.System.RouterPrefix)  // 公有路由 （无权限检查）
	privateGroup := engine.Group(config.ReportCfg.System.RouterPrefix) // 私有路由 （权限检查）
	{
		privateGroup.Use(middleware.JWTAuth()) //支持JWT
		//privateGroup.Use(middleware.CasbinHandler()) //权限管理
	}

	{

		initSystemRouter(publicGroup, privateGroup)

		initSysUserAuthRouter(publicGroup, privateGroup)  //用户授权
		initSysUserRouter(publicGroup, privateGroup)      //用户
		initOptionRecodeRouter(publicGroup, privateGroup) //记录操作

		// report
		initReportItemRouter(publicGroup, privateGroup) //报告api

	}
	zlog.Debug("load router success ")

	return nil
}
