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

var log = zlog.New("admin-router ")

func Init() error {

	var engine = gin.Default()

	if !config.AdminDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine.Use(middleware.CorsByRules())     // 按照配置的规则放行跨域请求
	engine.Use(middleware.GinRecovery(true)) // 处理异常

	{

		workPath, err := utils.GetWorkPath()
		if err != nil {
			return fmt.Errorf("os.Executable error " + err.Error())
		}
		// 应映页面
		utils.GinRouterView2Folder(engine, "/admin", workPath+"/views/admin")
		utils.GinRouterHome2Folder(engine, workPath+"/views/html")
	}

	publicGroup := engine.Group("/api")  // 公有路由 （无权限检查）
	privateGroup := engine.Group("/api") // 私有路由 （权限检查）
	{
		privateGroup.Use(middleware.JWTAuth()) //支持JWT
		//privateGroup.Use(middleware.CasbinHandler()) //权限管理
	}

	//
	initRouterSysSetup(publicGroup, privateGroup)  // 路由系统设置
	initRouterSysAuth(publicGroup, privateGroup)   // 路由授权
	initRouterSysAdmin(publicGroup, privateGroup)  // 路由管理员
	initRouterSysUser(publicGroup, privateGroup)   // 路由用户管理
	initRouterSysOption(publicGroup, privateGroup) // 路由操作记录

	go func() {
		log.Info("listen at ", config.AdminPort)
		engine.Run(fmt.Sprintf(":%d", config.AdminPort))
	}()
	time.Sleep(100 * time.Millisecond)

	return nil
}
