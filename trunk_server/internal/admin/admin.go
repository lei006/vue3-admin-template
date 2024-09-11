package admin

import (
	"vue3-admin-template/internal/admin/router"

	"github.com/lei006/zlog"
)

func Init() error {
	zlog.Debug("11111111")

	err := router.Init()
	if err != nil {
		return err
	}

	return nil
}
