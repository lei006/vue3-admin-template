package app

import (
	"vue3-admin-template/internal/admin"
	"vue3-admin-template/internal/admin/model"
	"vue3-admin-template/internal/config"
	"vue3-admin-template/pkg/daemon"

	"github.com/lei006/zlog"
)

func Init() error {
	// 设置工作目录
	WorkPath, err := daemon.GetWordPath()
	if err != nil {
		zlog.Error(err)
		return err
	}
	config.WorkPath = WorkPath

	// 初始化应用配置
	err = config.Init()
	if err != nil {
		zlog.Error(err)
		return err
	}

	//日志过滤
	for i := config.App.Zlog.Level; i < zlog.LogDebug; i++ {
		zlog.SetLogLevel(i + 1)
	}

	// 日志保存
	if config.App.Zlog.Save {
		zlog.SetSaveFile("logs.log", true)
	}

	// 强制控制台输出颜色
	if config.App.Zlog.Color {
		zlog.ForceConsoleColor()
	}

	err = model.Init(config.App.Model.DbType, config.App.Model.DbSource)
	if err != nil {
		zlog.Error(err)
		return err
	}

	// 初始化后台
	err = admin.Init()
	if err != nil {
		zlog.Error(err)
		return err
	}

	return nil
}
