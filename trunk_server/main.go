package main

import (
	"vue3-admin-template/internal/app"
	"vue3-admin-template/pkg/daemon"
	"vue3-admin-template/pkg/shell"

	"github.com/lei006/zlog"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {

	daemon.Run(main_run)

}

func main_run() {

	err := app.Init()
	if err != nil {
		zlog.Error("init app error:", err)
		return
	}

	// 保证 shell 程序一直在运行
	shell.RunUntilSignal()

}
