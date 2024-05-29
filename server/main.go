package main

import (
	"os"
	"yc-webreport-server/api"
	"yc-webreport-server/api/utils"
	"yc-webreport-server/config"

	"github.com/lei006/go-daemon/daemontool"
	"github.com/sohaha/zlsgo/zlog"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// 目录结构：
// controller  控制器
// service     服务层
// model       模型
// router      路由
// entity      结构体
// schedule    定时任务
// middleware  中间插件层
// extend      扩展--
// units       小组件

func main() {

	is_daemon := true

	err := config.OnInit()
	if err != nil {
		zlog.Error(err)
		return
	}

	///////////////////////////////////////////
	// 如果在vscode 则不是服务模式
	at_vscode, err := utils.RunAtVscode()
	if err != nil {
		zlog.Error("error: ")
		return
	}
	is_daemon = !at_vscode
	config.ReportCfg.RunAtVscode = at_vscode

	zlog.Debug("is_daemon ", is_daemon)

	///////////////////////////////////////////
	// 如果 参数 alone 则不是服务模式
	if len(os.Args) == 2 {
		if os.Args[1] == "alone" {
			is_daemon = false
		}
	}
	///////////////////////////////////////////
	// 如果 开始程序

	if is_daemon {
		//if true {
		daemonTool := daemontool.DefDaemonTool
		daemonTool.Run(config.AppName, config.AppDesc, RunApp)
	} else {
		RunApp()
	}

}

func RunApp() {
	//rtsp2web.Start()

	err := api.RunAndServer()
	if err != nil {
		zlog.Error(err)
	}

}
