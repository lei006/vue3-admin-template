package main

import (
	"vue3-admin-template/internal/exe_srv"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {

	exe_srv.Run()

}
