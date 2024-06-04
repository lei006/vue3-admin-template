package config

import (
	"fmt"
	"os"
	"path/filepath"
	"yc-webreport-server/utils"

	"github.com/sohaha/zlsgo/zlog"
	"github.com/spf13/viper"
)

var (
	ReportCfg     Server
	AppName       = "WebReport"
	AppDesc       = "工作站服务器"
	IsDaemon      = true
	DefConfigFile = "config.yaml"

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

type Server struct {
	RunAtVscode    bool   // 是否运行在vscode
	WorkPath       string // 工作路径
	ConfigFilePath string // 配置文件

	//
	Api Api `mapstructure:"api" json:"api" yaml:"api"`
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`

	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Logger  Logger  `mapstructure:"logger" json:"logger" yaml:"logger"`

	// gorm
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Sqlite Sqlite `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
}

func OnInit() error {

	///////////////////////////////////////////
	// 如果在vscode 则不是服务模式
	at_vscode, err := utils.RunAtVscode()
	if err != nil {
		zlog.Error("error: ")
		return err
	}

	ReportCfg.RunAtVscode = at_vscode

	if !ReportCfg.RunAtVscode {
		exePath, err := os.Executable()
		if err != nil {
			return fmt.Errorf("os.Executable error " + err.Error())
		}
		ReportCfg.WorkPath = filepath.Dir(exePath)
		ReportCfg.ConfigFilePath = ReportCfg.WorkPath + "\\" + DefConfigFile
	} else {

		pwdPath, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("os.Getwd error " + err.Error())
		}
		ReportCfg.WorkPath = pwdPath
		ReportCfg.ConfigFilePath = ReportCfg.WorkPath + "\\" + DefConfigFile
	}

	// 加载配置-配置
	err = loadConfig()
	if err != nil {
		return err
	}

	return nil
}

func loadConfig() error {

	///////////////////////////////////////////////
	// 1. 配置目录--支持 服务模式

	zlog.Debug("ReportCfg.RunAtVscode ", ReportCfg.RunAtVscode)
	zlog.Debug("config： ", ReportCfg.ConfigFilePath)

	v := viper.New()
	v.SetConfigFile(ReportCfg.ConfigFilePath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Fatal error config file: "+err.Error(), ReportCfg.ConfigFilePath)
	}

	if err = v.Unmarshal(&ReportCfg); err != nil {
		return fmt.Errorf("Fatal Unmarshal config file: " + err.Error())
	}

	return nil
}

func PrintInfo() {
	zlog.Info("")
	zlog.Info("")
	zlog.Info("")
	zlog.Info("---------------------------------------------------------------------------------------------")
	zlog.Info("----                                                                                   -----")
	zlog.Info("---------------------------------------------------------------------------------------------")
	zlog.Info("--AppName", AppName)
	zlog.Info("--AppDesc", AppDesc)
	zlog.Info("--RunAtVscode", ReportCfg.RunAtVscode)
	zlog.Info("--WorkPath", ReportCfg.WorkPath)
	zlog.Info("--ConfigFilePath", ReportCfg.ConfigFilePath)
	zlog.Info("--Logger.Path", ReportCfg.Logger.Path)
	zlog.Info("--Logger.SaveDay", ReportCfg.Logger.SaveDay)
	zlog.Info("")

}
