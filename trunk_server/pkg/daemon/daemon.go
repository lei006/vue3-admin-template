package daemon

import (
	"github.com/lei006/go-daemon/daemontool"
)

func Run(fun func()) {
	daemonTool := daemontool.DefDaemonTool
	ok, _ := daemontool.RunAtBuild()
	if ok {
		daemonTool.Run("test_app12", "desc 111测试333", fun)
	} else {
		fun()
	}

}

func GetWordPath() (string, error) {
	return daemontool.GetWordPath()
}
