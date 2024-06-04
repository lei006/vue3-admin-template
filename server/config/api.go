package config

type Api struct {
	Enable          bool   `mapstructure:"enable" json:"enable" yaml:"enable"`
	Debug           bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
	Addr            string `mapstructure:"addr" json:"addr" yaml:"addr"`                                        // 端口值
	UploadPrintPath string `mapstructure:"upload-print-path" json:"upload-print-path" yaml:"upload-print-path"` // 上传文件路径
}
