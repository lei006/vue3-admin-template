package app

import (
	"fmt"
	"vue3-admin-template/internal/admin"
	"vue3-admin-template/internal/config"
	"vue3-admin-template/pkg/daemon"

	"github.com/lei006/zlog"
)

var log = zlog.New("app")

func Init() error {

	WorkPath, err := daemon.GetWordPath()
	if err != nil {
		fmt.Println("get word path error:", err)
		return err
	}

	config.WorkPath = WorkPath

	err = admin.Init()
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
