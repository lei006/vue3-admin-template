package main

import (
	"fmt"
	"time"
	"vue3-admin-template/internal/app"
	"vue3-admin-template/internal/config"
	"vue3-admin-template/pkg/daemon"
	"vue3-admin-template/pkg/shell"

	"github.com/lei006/zlog"
)

var log = zlog.New("main ")

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
		fmt.Println("init app error:", err)
		return
	}

	for {

		log.Debug("WorkPath =", config.WorkPath)
		time.Sleep(time.Second)
	}

	// 保证 shell 程序一直在运行
	shell.RunUntilSignal()

}
