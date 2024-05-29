package captcha

import (
	"yc-webreport-server/api/utils"
	"yc-webreport-server/config"

	"github.com/songzhibin97/gkit/cache/local_cache"
)

var (
	BlackCache local_cache.Cache
)

func LoadConfig() error {

	dr, err := utils.ParseDuration(config.ReportCfg.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(config.ReportCfg.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)

	return nil
}
