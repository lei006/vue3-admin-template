package config

import (
	"flag"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/adapter/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/lei006/go-assist/utils"
)

var (
	AppIssuer     = "徐州云宝电子科技"
	HttpPort      = int(8180)
	DebugMode     = 0
	LogEnable     = 1
	TokenKeepTime = int(365 * 24 * 60) //分钟

	Debug = false

	WorkPath = "./"

	DbDataSource = ""
	KeyString    = "eyJwdWJfa2V5IjoiLS0tLS1CRUdJTiBlY2MgcHVibGljIGtleS0tLS0tXG5NSUdiTUJBR0J5cUdTTTQ5QWdFR0JTdUJCQUFqQTRHR0FBUUFuYURkYUNmN0hvMmNmWk4xSW5xTURxcloyNW4vXG5LUmNJYldSdlgxbTdiRFYzQm5mVjFVZHhJQkg2WTBxdjExcGN1QmxsL1cvVG43cHFBU2U3a3ZnZElyb0JFbHlTXG45bnphbjdNemRxY0V3Wkt1bFgyVDlJWlBMWngrdlpxVFREZUZ1K2VqcndyQ28yMStuRGNiVm5IQ01jaE1Vb2d0XG4rVG5sNWZmdmlVc0QwbVNVTVg0PVxuLS0tLS1FTkQgZWNjIHB1YmxpYyBrZXktLS0tLVxuIiwicHJpX2tleSI6Ii0tLS0tQkVHSU4gZWNjIHByaXZhdGUga2V5LS0tLS1cbk1JSGNBZ0VCQkVJQXVXRkplYkQrQWJXd09FditPTHpMUUFUYzVUQUhReXlBM2g3MStJd0IrQithd3R5cUdVSkZcbnNDYmU4UkNPV2RmMDNpbVVhL3Azc2Z0L1YrWlNMYVhReHY2Z0J3WUZLNEVFQUNPaGdZa0RnWVlBQkFDZG9OMW9cbkovc2VqWng5azNVaWVvd09xdG5ibWY4cEZ3aHRaRzlmV2J0c05YY0dkOVhWUjNFZ0VmcGpTcS9YV2x5NEdXWDlcbmI5T2Z1bW9CSjd1UytCMGl1Z0VTWEpMMmZOcWZzek4ycHdUQmtxNlZmWlAwaGs4dG5INjltcE5NTjRXNzU2T3ZcbkNzS2piWDZjTnh0V2NjSXh5RXhTaUMzNU9lWGw5KytKU3dQU1pKUXhmZz09XG4tLS0tLUVORCBlY2MgcHJpdmF0ZSBrZXktLS0tLVxuIn0="
)

func init() {

	ReLoadConfig()

	ConfigLogger()
}

func ConfigLogger() {
	if utils.Is_RunAtVs() {
		//运行在vs 中
		WorkPath = utils.GetWorkPath() + "/"
	} else {
		WorkPath = utils.GetBinPath() + "/"
	}

	//////////////////////////////////////////////////////////
	// 配置运行路径 -- 不在 vs
	if Debug == false && !utils.Is_RunAtVs() {
		WorkPath = utils.GetBinPath() + "/"
	}

	//////////////////////////////////////////////////////////
	// 配置日志...
	if LogEnable == 1 {
		log_filename := WorkPath + strings.ToLower(AppName) + ".log"
		logs.SetLogger(logs.AdapterFile, `{"filename":"`+log_filename+`", "level":7}`)
	}

}

func ReLoadConfig() error {

	configpath := flag.String("c", "conf/app.conf", "The default to load the conf/app.conf")
	flag.Parse()

	config_filenamepath := *configpath
	//config_filenamepath = strings.Replace(config_filenamepath, "/", "\\", -1) //将/替换成\\
	//config_filenamepath = "conf/app.conf"

	iniconf, err := config.NewConfig("ini", config_filenamepath)
	if err != nil {
		fmt.Println("read config file error: ", err.Error())
		return err
	}

	AppName = iniconf.DefaultString("appname", AppName)
	HttpPort = iniconf.DefaultInt("http_port", HttpPort)
	DebugMode = iniconf.DefaultInt("debug_mode", DebugMode)
	TokenKeepTime = iniconf.DefaultInt("TokenKeepTime", TokenKeepTime)
	DbDataSource = iniconf.DefaultString("DbDataSource", DbDataSource)
	LogEnable = iniconf.DefaultInt("LogEnable", LogEnable)
	KeyString = iniconf.DefaultString("key", KeyString)

	return nil
}
