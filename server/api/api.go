package api

import (
	"yc-webreport-server/api/extend/captcha"
	"yc-webreport-server/api/extend/database"
	"yc-webreport-server/api/extend/logger"
	"yc-webreport-server/api/extend/webserver"
	"yc-webreport-server/api/model"
	"yc-webreport-server/config"
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

	// 加载配置-配置
	err := config.LoadConfig()

	if err != nil {
		return err
	}

	// 加载配置-日志系统
	err = logger.LoadConfig()
	if err != nil {
		return err
	}

	config.PrintInfo()

	// 加载配置-验证码
	err = captcha.LoadConfig()
	if err != nil {
		return err
	}

	// 加载配置-数据库
	err = database.LoadConfig()
	if err != nil {
		return err
	}

	err = model.OnInit()
	if err != nil {
		return err
	}

	// 开始运行
	return webserver.RunAndServer()

}
