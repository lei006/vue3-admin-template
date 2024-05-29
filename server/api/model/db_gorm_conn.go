package model

import (
	"log"
	"os"
	"time"
	"yc-webreport-server/config"

	"github.com/sohaha/zlsgo/zlog"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	g_db *gorm.DB //数据库
)

func OnInit() error {

	dbtype := config.ReportCfg.System.DbType
	zlog.Debug(config.ReportCfg.System)
	zlog.Debug("LoadConfig database : " + dbtype)

	db, err := getGorm(dbtype)
	if err != nil {
		zlog.Error("get gorm error: " + err.Error())
		return err
	}

	g_db = db

	registerTables()

	return nil
}

func getGorm(dbType string) (*gorm.DB, error) {
	switch dbType {
	case "mysql":
		return gormMysql()
		/*
			case "pgsql":
				return GormPgSql()
			case "oracle":
				return GormOracle()
			case "mssql":
				return GormMssql()
		*/
	case "sqlite":
		return gormSqlite()
	default:
		return gormMysql()
	}
}

func registerTables() {
	err := g_db.AutoMigrate(
		//SysApi{},
		SysUser{},
		JwtBlacklist{},
		SysOperationRecord{},
	)

	if err != nil {
		zlog.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	zlog.Info("register table success")
}

type DBBASE interface {
	GetLogMode() string
}

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	gorm_config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	_default := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	var logMode DBBASE
	switch config.ReportCfg.System.DbType {
	case "mysql":
		logMode = &config.ReportCfg.Mysql
		/*
			case "pgsql":
				logMode = &config.ReportCfg.Pgsql
			case "oracle":
				logMode = &config.ReportCfg.Oracle
		*/
	default:
		logMode = &config.ReportCfg.Mysql
	}

	switch logMode.GetLogMode() {
	case "silent", "Silent":
		gorm_config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		gorm_config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		gorm_config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		gorm_config.Logger = _default.LogMode(logger.Info)
	default:
		gorm_config.Logger = _default.LogMode(logger.Info)
	}
	return gorm_config
}
