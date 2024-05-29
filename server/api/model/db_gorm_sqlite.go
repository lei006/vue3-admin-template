package model

import (
	"errors"
	"yc-webreport-server/config"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// GormSqlite 初始化Sqlite数据库
func gormSqlite() (*gorm.DB, error) {
	s := config.ReportCfg.Sqlite

	if s.Dbname == "" {
		return nil, errors.New("not config sqlite")
	}

	if db, err := gorm.Open(sqlite.Open(s.Dsn()), Gorm.Config(s.Prefix, s.Singular)); err != nil {
		return nil, errors.New("not config sqlite")
	} else {
		sqlDB, err := db.DB()
		if err != nil {
			return nil, err
		}

		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)

		return db, nil
	}
}

// GormSqliteByConfig 初始化Sqlite数据库用过传入配置
func gormSqliteByConfig(s config.Sqlite) *gorm.DB {
	if s.Dbname == "" {
		return nil
	}

	if db, err := gorm.Open(sqlite.Open(s.Dsn()), Gorm.Config(s.Prefix, s.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		return db
	}
}
