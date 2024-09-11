package model

import (
	"fmt"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/lei006/zlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	g_db *gorm.DB //数据库

	//DbType   = "sqlite"  // sqlite or mysql
	//DbSource = "data.db" // 实际使用的数据库

	//DbType   = "mysql" // sqlite or mysql
	//DbSource = "root:wLei6700413@@tcp(127.0.0.1:3306)/media_nvr?charset=utf8mb4&parseTime=True&loc=Local"

	//这是Mysql示例:  DbSource = user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	//这是sqlite示例:  DbSource = data.db
)

type BASE_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"id"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

func Init(dbType string, dbSource string) error {
	zlog.Debug("111111111")
	g_db = getdb(dbType, dbSource)

	return nil
}

func getdb(dbType string, dbSource string) *gorm.DB {

	if g_db == nil {

		var tmp_db *gorm.DB
		///////////////////////////////////////
		// 如果没有连接，则连接库
		if dbType == "mysql" {

			mysqlConfig := mysql.Config{
				DSN:                       dbSource, // DSN data source name
				DefaultStringSize:         191,      // string 类型字段的默认长度
				SkipInitializeWithVersion: false,    // 根据版本自动配置
			}
			gorm_config := &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					//TablePrefix:   "",
					SingularTable: false,
				},
				DisableForeignKeyConstraintWhenMigrating: true,
			}
			db, err := gorm.Open(mysql.New(mysqlConfig), gorm_config)
			if err != nil {
				zlog.Fatal(err)
				panic(fmt.Sprintf("Failed to connect to database: %v", err))
			}

			tmp_db = db
		} else if dbType == "sqlite" {
			db, err := gorm.Open(sqlite.Open(dbSource), &gorm.Config{})
			if err != nil {
				zlog.Fatal(err)
				panic(fmt.Sprintf("Failed to connect to database: %v", err))
			}
			tmp_db = db
		} else {
			zlog.Error("unknow database type " + dbType)
			panic("unknow database type " + dbType)
		}

		///////////////////////////////////////
		// 连接成功，自动生成表
		err := tmp_db.AutoMigrate(
			SysUser{},
			SysSetup{},
			JwtBlacklist{},
		)
		if err != nil {
			zlog.Error(err.Error())
			panic("register table failed" + err.Error())
		}
		g_db = tmp_db
	}

	return g_db
}

type DBBASE interface {
	GetLogMode() string
}

func ErrRecordNotFound(err error) bool {
	if err == gorm.ErrRecordNotFound {
		return true
	}
	return false
}

/*
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

	return gorm_config
}
*/
