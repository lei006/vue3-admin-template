package router

import (
	"fmt"
	"time"
	"vue3-admin-template/internal/admin/middleware"
	"vue3-admin-template/internal/config"
	"vue3-admin-template/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/lei006/zlog"
)

func Init() error {

	var engine = gin.Default()

	if !config.App.Admin.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine.Use(middleware.LimitIp())         // 限制IP
	engine.Use(middleware.CorsByRules())     // 跨域请求
	engine.Use(middleware.GinRecovery(true)) // 处理异常
	engine.Use(middleware.RecordOption())    // 记录操作

	{

		workPath, err := utils.GetWorkPath()
		if err != nil {
			return fmt.Errorf("os.Executable error " + err.Error())
		}
		// 应映页面
		utils.GinRouterView2Folder(engine, "/admin", workPath+"/views/admin")
		utils.GinRouterHome2Folder(engine, workPath+"/views/html")

		zlog.Info("映射html: ", "/     ", "------>", workPath+"/views/html")
		zlog.Info("映射目录: ", "/admin", "------>", workPath+"/views/admin")

	}

	publicGroup := engine.Group("/api")  // 公有路由 （无权限检查）
	privateGroup := engine.Group("/api") // 私有路由 （权限检查）
	{
		privateGroup.Use(middleware.JWTAuth()) //支持JWT
		//privateGroup.Use(middleware.CasbinHandler()) //权限管理
	}

	//
	initRouterSysAuthUser(publicGroup, privateGroup)  // 认证模块
	initRouterSysAuthAdmin(publicGroup, privateGroup) // 认证模块
	initRouterSysAdmin(publicGroup, privateGroup)     // 路由管理员
	initRouterSysUser(publicGroup, privateGroup)      // 路由用户管理
	initRouterSysOption(publicGroup, privateGroup)    // 路由操作记录
	initRouterSysSetup(publicGroup, privateGroup)     // 路由系统设置
	initRouterSysLimitIp(publicGroup, privateGroup)   // 路由系统设置
	initRouterSysLicense(publicGroup, privateGroup)   // 路由系统设置
	initRouterSysAbout(publicGroup, privateGroup)     // 路由系统设置

	go func() {
		zlog.Info("Admin listen at ", config.App.Admin.Port)
		engine.Run(fmt.Sprintf(":%d", config.App.Admin.Port))
	}()
	time.Sleep(100 * time.Millisecond)

	return nil
}
