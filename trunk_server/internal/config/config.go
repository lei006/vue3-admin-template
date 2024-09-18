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
	App            AppConfig
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

	// 数据库初始化
	err = db_model.Init(App.Model.DbType, App.Model.DbSource)
	if err != nil {
		zlog.Error(err)
		return err
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

	////////////////////////////////////////////////////
	// 检查预置数据
	err = checkPreData()
	if err != nil {
		zlog.Error(err)
		return err
	}

	return nil
}

func checkPreData() error {
	modelAdmin := db_model.SysAdmin{}
	_, err := modelAdmin.FindOrCreate("admin", App.Admin.Password)
	if err != nil {
		zlog.Error(err)
		return err
	}

	return nil
}
