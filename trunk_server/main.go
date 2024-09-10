package main

import (
	"fmt"
	"time"
	"vue3-admin-template/pkg/daemon"
	"vue3-admin-template/pkg/shell"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {

	daemon.Run(main_run)

}

func main_run() {

	for true {

		path, err := daemon.GetWordPath()
		if err != nil {
			fmt.Println("get word path error:", err)
			continue
		}

		fmt.Println("xxxxxxx", path)
		time.Sleep(time.Second)
	}

	// 保证 shell 程序一直在运行
	shell.RunUntilSignal()

}
