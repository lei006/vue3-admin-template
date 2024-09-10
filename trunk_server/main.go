package main

import (
	"fmt"
	"time"

	"github.com/lei006/go-daemon/daemontool"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {

	daemonTool := daemontool.DefDaemonTool

	ok, _ := daemontool.RunAtBuild()
	if ok {
		daemonTool.Run("test_app12", "desc 111测试333", Run)
	} else {
		Run()
	}
}

func Run() {

	for true {
		fmt.Println("xxxxxxx", daemontool.DefDaemonTool.GetWordPath())
		time.Sleep(time.Second)
	}

}
