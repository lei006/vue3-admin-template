package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sohaha/zlsgo/zlog"
	"github.com/spf13/viper"
)

var (
	ReportCfg Server
	AppName   = "WebReport"
	AppDesc   = "工作站服务器"

	DefConfigFile = "config.yaml"
)

type Server struct {
	RunAtVscode    bool   // 是否运行在vscode
	WorkPath       string // 工作路径
	ConfigFilePath string // 配置文件

	//
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`

	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Logger  Logger  `mapstructure:"logger" json:"logger" yaml:"logger"`

	// gorm
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Sqlite Sqlite `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}

func OnInit() error {
	return nil
}

func LoadConfig() error {

	///////////////////////////////////////////////
	// 1. 配置目录--支持 服务模式

	configFile := DefConfigFile
	if !ReportCfg.RunAtVscode {
		exePath, err := os.Executable()
		if err != nil {
			return fmt.Errorf("os.Executable error " + err.Error())
		}
		ReportCfg.WorkPath = filepath.Dir(exePath)
		ReportCfg.ConfigFilePath = ReportCfg.WorkPath + "/" + configFile
	} else {

		pwdPath, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("os.Getwd error " + err.Error())
		}
		ReportCfg.WorkPath = pwdPath
		ReportCfg.ConfigFilePath = ReportCfg.WorkPath + "/" + configFile
	}

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
