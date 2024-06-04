package api

import (
	"yc-webreport-server/api/router"
	"yc-webreport-server/config"

	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
)

// 目录结构：
// controller  控制器
// service     服务层
// model       模型
// router      路由
// schedule    定时任务
// middleware  中间插件层
// extend      扩展层
// units       小组件

func RunAndServer() error {

	if !config.ReportCfg.Api.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()

	// 加载-路由
	err := router.LoadRouter(engine)
	if err != nil {
		return err
	}

	addr := config.ReportCfg.Api.Addr
	zlog.Info("WebServer Listen :" + addr)
	err = engine.Run(addr)
	if err != nil {
		zlog.Debug("WebServer error :" + err.Error())
	}

	return nil
}
