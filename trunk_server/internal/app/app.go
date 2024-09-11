package app

import (
	"vue3-admin-template/internal/admin"
	"vue3-admin-template/internal/config"
	"vue3-admin-template/pkg/daemon"

	"github.com/lei006/zlog"
)

func Init() error {

	//日志保存
	if config.LogSave {
		zlog.SetSaveFile("logs.log", true)
	} else {
		// 强制控制台输出颜色
		zlog.ForceConsoleColor()
	}

	// 设置工作目录
	WorkPath, err := daemon.GetWordPath()
	if err != nil {
		zlog.Error(err)
		return err
	}
	config.WorkPath = WorkPath

	// 初始化后台
	err = admin.Init()
	if err != nil {
		zlog.Error(err)
		return err
	}

	return nil
}
