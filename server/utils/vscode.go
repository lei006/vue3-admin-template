package utils

import (
	"fmt"
	"os"
	"strings"
)

func RunAtVscode() (bool, error) {

	exePath, err := os.Executable()
	if err != nil {
		return false, fmt.Errorf("os.Executable error " + err.Error())
	}

	contains := strings.Contains(exePath, "go-build")
	return contains, nil
}
