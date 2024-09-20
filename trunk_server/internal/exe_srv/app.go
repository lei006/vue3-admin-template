package exe_srv

import (
	"vue3-admin-template/internal/admin"
	"vue3-admin-template/internal/config"
	"vue3-admin-template/internal/db_model"
	"vue3-admin-template/internal/license"
	"vue3-admin-template/pkg/daemon"
	"vue3-admin-template/pkg/shell"

	"github.com/lei006/zlog"
)

func Run() error {

	daemon.Run(daemon_run)

	return nil
}

func daemon_run() {
	err := srv_run()
	if err != nil {
		zlog.Error("run server error:", err)
		return
	}

	shell.RunUntilSignal()
}

func srv_run() error {

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

	// 加载授权文件
	err = license.Init()
	if err != nil {
		zlog.Error("load License error:", err)
		return err
	}

	// 初始化应用配置
	err = config.Init()
	if err != nil {
		zlog.Error(err)
		return err
	}

	// 数据库初始化
	err = db_model.Init(config.App.Model.DbType, config.App.Model.DbSource)
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

	////////////////////////////////////////////////////
	// 预置数据
	modelAdmin := db_model.SysAdmin{}
	_, err = modelAdmin.FindOrCreate("admin", config.App.Admin.Password)
	if err != nil {
		zlog.Error(err)
		return err
	}
	return nil
}
