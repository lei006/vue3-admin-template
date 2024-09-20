package exe_app

import (
	"github.com/lei006/zlog"
)

func Init() error {

	zlog.ForceConsoleColor()

	zlog.Info("=============================================================")
	zlog.Info("=============================================================")
	zlog.Info("=========      欢迎使用 vue3-admin-template            =======")
	zlog.Info("-------------------------------------------------------------")
	zlog.Info("")

	return nil
}
