package model

import (
	"errors"
)

// 配置文件结构体
type SysSetup struct {
	BASE_MODEL
	Name string `json:"name" gorm:"column:name;uniqueIndex;comment:名称"`
	Data string `json:"data"  gorm:"column:data;type:longtext;comment:数值"`
	Desc string `json:"desc" gorm:"column:desc;type:tinytext;comment:描述"`
}

func (model *SysSetup) TableName() string {
	return "sys_setup"
}

func (model *SysSetup) PatchOne(id string, field string, data interface{}) error {

	result := g_db.Model(&SysSetup{}).Where("id = ?", id).Update(field, data)
	if result.RowsAffected == 0 {
		return errors.New("no rows were updated")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (model *SysSetup) GetOne(id string) (*SysSetup, error) {

	retVal := &SysSetup{}
	err := g_db.Where("id = ?", id).First(&retVal).Error
	return retVal, err
}

func (model *SysSetup) GetOneByName(name string) (*SysSetup, error) {
	// 通过username 取得一行
	retVal := &SysSetup{}
	err := g_db.Where("name = ?", name).First(retVal).Error
	return retVal, err
}

func (model *SysSetup) GetList() (list []SysSetup, total int64, err error) {

	// 创建db
	db := g_db.Model(&SysSetup{})
	var items []SysSetup
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Find(&items).Error
	return items, total, err
}
