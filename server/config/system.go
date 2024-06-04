package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                   // 环境值
	GinDebug      bool   `mapstructure:"gin-debug" json:"gin-debug" yaml:"gin-debug"` // gin-模式
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`       // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`    // Oss类型
	RouterPrefix  string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
}

type Logger struct {
	Path    string `mapstructure:"path" json:"path" yaml:"path"`             // 环境值
	SaveDay int    `mapstructure:"save_day" json:"save_day" yaml:"save_day"` // 保存日期
	Enable  bool   `mapstructure:"enable" json:"enable" yaml:"enable"`       // 使用redis
}
