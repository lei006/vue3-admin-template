package app

import (
	"vue3-admin-template/internal/admin"
	"vue3-admin-template/internal/config"
	"vue3-admin-template/internal/db_model"
	"vue3-admin-template/pkg/daemon"

	"github.com/lei006/zlog"
)

func Init() error {

	zlog.ForceConsoleColor()

	zlog.Info("=============================================================")
	zlog.Info("=============================================================")
	zlog.Info("=========      欢迎使用 vue3-admin-template            =======")
	zlog.Info("-------------------------------------------------------------")
	zlog.Info("")

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

	db_model.FindOrCreateAbout("程序名称", config.AppName, "程序描述")

	// 初始化后台
	err = admin.Init()
	if err != nil {
		zlog.Error(err)
		return err
	}

	return nil
}
