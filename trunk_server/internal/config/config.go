package config

import (
	"fmt"

	"github.com/lei006/zlog"
	"github.com/spf13/viper"
)

var (
	WorkPath = "./" // 工作路径

	ConfigFileName = "config.yaml"
	App            AppConfig
)

type ModelConfig struct {
	DbType   string `mapstructure:"dbtype" json:"dbtype" yaml:"dbtype"`
	DbSource string `mapstructure:"dbsource" json:"dbsource" yaml:"dbsource"`
	Test     string `mapstructure:"test" json:"test" yaml:"test"`
}

type AdminConfig struct {
	Debug bool `mapstructure:"debug" json:"debug" yaml:"debug"`
	Port  int  `mapstructure:"port" json:"port" yaml:"port"`
}

type LogConfig struct {
	Save  bool `mapstructure:"save" json:"save" yaml:"save"`
	Level int  `mapstructure:"level" json:"level" yaml:"level"`
	Color bool `mapstructure:"color" json:"color" yaml:"color"`
}

type AppConfig struct {
	Admin AdminConfig `mapstructure:"admin" json:"admin" yaml:"admin"`
	Zlog  LogConfig   `mapstructure:"zlog" json:"zlog" yaml:"zlog"`
	Model ModelConfig `mapstructure:"model" json:"model" yaml:"model"`
}

func Init() error {

	ConfigFilePath := fmt.Sprintf("%s/%s", WorkPath, ConfigFileName)

	v := viper.New()
	v.SetConfigFile(ConfigFilePath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Fatal error config file: "+err.Error(), ConfigFilePath)
	}

	if err = v.Unmarshal(&App); err != nil {
		return fmt.Errorf("Fatal Unmarshal config file: " + err.Error())
	}

	// 如果 zlog.level 大于8，则设置为8
	if App.Zlog.Level > zlog.LogDebug || App.Zlog.Level < zlog.LogFatal {
		App.Zlog.Level = zlog.LogDebug
	}

	//如果 admin.port 小于等于0 大于65535，则设置为8090
	if App.Admin.Port <= 0 || App.Admin.Port > 65535 {
		App.Admin.Port = 8090
	}

	return nil
}
