package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/sohaha/zlsgo/zlog"
)

func RunAtVscode() (bool, error) {

	exePath, err := os.Executable()
	if err != nil {
		return false, fmt.Errorf("os.Executable error " + err.Error())
	}

	contains := strings.Contains(exePath, "go-build")

	zlog.Debug(exePath, contains)

	return contains, nil
}
