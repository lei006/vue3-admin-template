package admin

import (
	"vue3-admin-template/internal/admin/router"
)

func Init() error {

	err := router.Init()
	if err != nil {
		return err
	}

	return nil
}
