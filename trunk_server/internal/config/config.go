package config

import (
	"fmt"
	"vue3-admin-template/internal/db_model"

	"github.com/lei006/zlog"
	"github.com/spf13/viper"
)

var (
	WorkPath = "./" // 工作路径

	ConfigFileName = "config.yaml"
	AppName        = "我是应用的程序名"
	App            AppConfig
	HardSn         string
	LicenseFile    = "license.lic"

	LicenseCheck bool //授权检查是否成功
	Lic          db_model.LicenseStruct

	LicensePubKey = "MIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQAW9CZ4zNrhuDsmNSKst2T06S62QTO02B12GXMbWPrbUOApbIkO0Sswc3n45bCRDHp0fWyMMHUF/GLS12hurRyElYBXEQQLRxPl/VsXg+FLO+KkhIxVtaIhuYgsnrr3Fa0HZCa/7CtrfrjnlrzJD4tNJUjdA9gPP0GJCioOw32pH3PypU="
	LicensePriKey = "MIHcAgEBBEIB4ZuxKV3TGWxV3rWeN4khLfPLXfcU+cTUyIRzZ2oh2pHEahbetiNJK3rCKF0lHkeLA7nhvQkCUiS1TCpqMqc0IamgBwYFK4EEACOhgYkDgYYABABb0JnjM2uG4OyY1Iqy3ZPTpLrZBM7TYHXYZcxtY+ttQ4ClsiQ7RKzBzefjlsJEMenR9bIwwdQX8YtLXaG6tHISVgFcRBAtHE+X9WxeD4Us74qSEjFW1oiG5iCyeuvcVrQdkJr/sK2t+uOeWvMkPi00lSN0D2A8/QYkKKg7Dfakfc/KlQ=="
	//pubkey: MIGbMBAGByqGSM49AgEGBSuBBAAjA4GGAAQAW9CZ4zNrhuDsmNSKst2T06S62QTO02B12GXMbWPrbUOApbIkO0Sswc3n45bCRDHp0fWyMMHUF/GLS12hurRyElYBXEQQLRxPl/VsXg+FLO+KkhIxVtaIhuYgsnrr3Fa0HZCa/7CtrfrjnlrzJD4tNJUjdA9gPP0GJCioOw32pH3PypU=
	//prikey: MIHcAgEBBEIB4ZuxKV3TGWxV3rWeN4khLfPLXfcU+cTUyIRzZ2oh2pHEahbetiNJK3rCKF0lHkeLA7nhvQkCUiS1TCpqMqc0IamgBwYFK4EEACOhgYkDgYYABABb0JnjM2uG4OyY1Iqy3ZPTpLrZBM7TYHXYZcxtY+ttQ4ClsiQ7RKzBzefjlsJEMenR9bIwwdQX8YtLXaG6tHISVgFcRBAtHE+X9WxeD4Us74qSEjFW1oiG5iCyeuvcVrQdkJr/sK2t+uOeWvMkPi00lSN0D2A8/QYkKKg7Dfakfc/KlQ==

)

type ModelConfig struct {
	DbType   string `mapstructure:"dbtype" json:"dbtype" yaml:"dbtype"`
	DbSource string `mapstructure:"dbsource" json:"dbsource" yaml:"dbsource"`
}

type AdminConfig struct {
	Debug    bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type LogConfig struct {
	SaveDay int  `mapstructure:"saveday" json:"saveday" yaml:"saveday"`
	Level   int  `mapstructure:"level" json:"level" yaml:"level"`
	Color   bool `mapstructure:"color" json:"color" yaml:"color"`
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

	////////////////////////////////////////////////////
	// 配置检查
	//如果 admin.port 小于等于0 大于65535，则设置为8090
	if App.Admin.Port <= 0 || App.Admin.Port > 65535 {
		zlog.Warn("admin.port is invalid, set to 8090")
		App.Admin.Port = 8090
	}

	// 如果 zlog.level 大于8，则设置为8
	if App.Zlog.Level > zlog.LogDebug || App.Zlog.Level < zlog.LogFatal {
		zlog.Warn("zlog.level is invalid, set to 8")
		App.Zlog.Level = zlog.LogDebug
	}

	////////////////////////////////////////////////////
	// 配置使用
	//日志过滤
	for i := App.Zlog.Level; i < zlog.LogDebug; i++ {
		zlog.SetLogLevel(i + 1)
	}

	// 日志保存
	if App.Zlog.SaveDay > 0 {
		zlog.SetSaveFile("logs.log", true)
		zlog.LogMaxDurationDate = App.Zlog.SaveDay
	}

	// 强制控制台输出颜色
	zlog.ForceConsoleColor()

	return nil
}
