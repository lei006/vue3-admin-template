// 自动生成模板SysOperationRecord
package db_model

import (
	"errors"

	"github.com/sohaha/zlsgo/zlog"
)

// 如果含有time.Time 请自行import time包
type SysOption struct {
	BASE_MODEL
	FromIp    string `json:"fromip" form:"fromip" gorm:"column:fromip;comment:请求ip"`      // 请求ip
	Typ       string `json:"typ" form:"typ" gorm:"column:typ;comment:type"`               // 0
	MsgText01 string `json:"msgtext01" form:"msgtext01" gorm:"column:msgtext01;comment:"` // 0
	MsgText02 string `json:"msgtext02" form:"msgtext02" gorm:"column:msgtext02;comment:"` // 0
	MsgText03 string `json:"msgtext03" form:"msgtext03" gorm:"column:msgtext03;comment:"` // 0
	MsgText04 string `json:"msgtext04" form:"msgtext04" gorm:"column:msgtext04;comment:"` // 0
}

func (SysOption) TableName() string {
	return "sys_option"
}

func (model *SysOption) Create(newVal *SysOption) (err error) {

	result := g_db.Create(newVal)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no rows were inserted")
	}

	return nil
}

func (model *SysOption) DeleteMany(ids []int) (err error) {
	result := g_db.Unscoped().Delete(&[]SysOption{}, "id in ?", ids)
	if result.Error != nil {
		return result.Error
	}

	return err
}

func (model *SysOption) GetPage(page PageInfo) (list []SysOption, total int64, err error) {

	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)

	// 创建db
	db := g_db.Model(&SysOption{})

	if len(page.Keyword) > 0 {
		db = db.Where("typ LIKE ?", "%"+page.Keyword+"%").Or("fromip LIKE ?", "%"+page.Keyword+"%").Or("msgtext01 LIKE ?", "%"+page.Keyword+"%").Or("msgtext02 LIKE ?", "%"+page.Keyword+"%").Or("msgtext03 LIKE ?", "%"+page.Keyword+"%").Or("msgtext04 LIKE ?", "%"+page.Keyword+"%")
	}

	var items []SysOption
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	result := db.Limit(limit).Offset(offset).Order("id desc").Find(&items)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return items, total, err
}

func SysOptionLog(typ, msg1, msg2, msg3, msg4, FromIp string) error {

	item := SysOption{
		FromIp:    FromIp,
		Typ:       typ,
		MsgText01: msg1,
		MsgText02: msg2,
		MsgText03: msg3,
		MsgText04: msg4,
	}

	var sys_options SysOption
	err := sys_options.Create(&item)
	if err != nil {
		zlog.Error("SysOptionLog", err)
		return err
	}

	return nil
}
