package webserver

import (
	"fmt"

	"yc-webreport-server/api/router"
	"yc-webreport-server/config"

	"github.com/gin-gonic/gin"
	"github.com/sohaha/zlsgo/zlog"
)

func RunAndServer() error {

	engine := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	// 加载-路由
	err := router.LoadRouter(engine)
	if err != nil {
		return err
	}

	addr := fmt.Sprintf(":%d", config.ReportCfg.System.Addr)
	zlog.Info("WebServer Listen :" + addr)
	err = engine.Run(addr)
	if err != nil {
		zlog.Debug("WebServer error :" + err.Error())
	}

	return nil
}
