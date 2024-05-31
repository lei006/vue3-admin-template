package logger

import (
	"yc-webreport-server/config"

	"github.com/sohaha/zlsgo/zlog"
)

func OnInit() error {

	logPath := config.ReportCfg.Logger.Path
	if logPath == "" {
		config.ReportCfg.Logger.Path = config.ReportCfg.WorkPath + "\\logs.log"
	}
	zlog.SetSaveFile(config.ReportCfg.Logger.Path, true)
	zlog.LogMaxDurationDate = config.ReportCfg.Logger.SaveDay
	return nil
}
