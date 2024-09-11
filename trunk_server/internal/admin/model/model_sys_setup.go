package model

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/lei006/zlog"
	"gorm.io/gorm"
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

func (model *SysSetup) Create(val *SysSetup) error {
	err := g_db.Create(val).Error
	return err
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
	if err != nil {
		return nil, err
	}

	return retVal, nil
}

func (model *SysSetup) PatchOneByName(name string, data interface{}) error {

	result := g_db.Model(&SysSetup{}).Where("name = ?", name).Update("data", data)
	if result.RowsAffected == 0 {
		return errors.New("no rows were updated")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (model *SysSetup) IsExistByName(name string) (*SysSetup, error) {
	// 通过username 取得一行
	retVal := &SysSetup{}
	ret := g_db.Where("name = ?", name).First(retVal)
	// 判断是否存在
	if ret.RowsAffected == 0 {
		return nil, ret.Error
	}
	return retVal, nil
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

func (model *SysSetup) NoFoundCreate(name string, data string, desc string) error {

	val := SysSetup{
		Name: name,
		Data: data,
		Desc: desc,
	}
	err := g_db.Where("name = ?", name).First(&val).Error

	if err == nil {
		return nil
	}

	if err != gorm.ErrRecordNotFound { // 如果不是未找到错误，则说明用户已存在
		return err
	}

	err = g_db.Create(&val).Error
	if err != nil {
		zlog.Error(err)
		return err
	}
	return nil
}

func (model *SysSetup) AutoGetStringByName(name string) (string, error) {
	ret, err := model.GetOneByName(name)

	if err != nil {
		zlog.Error(err)
		return "", err
	}

	return ret.Data, nil
}

func (model *SysSetup) AutoSetStringByName(name string, data string) error {

	err := model.PatchOneByName(name, data)

	if err != nil {
		zlog.Error(err)
		return err
	}

	return nil
}

func (model *SysSetup) AutoGetIntByName(name string) (int, error) {

	val_str, err := model.AutoGetStringByName(name)
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(val_str)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (model *SysSetup) AutoSetIntByName(name string, data int) error {
	return model.AutoSetStringByName(name, fmt.Sprintf("%d", data))
}
