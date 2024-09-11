package model

import (
	"errors"

	"gorm.io/gorm"
)

type MediaOption struct {
	BASE_MODEL
	Name string `json:"name" gorm:"column:name;uniqueIndex;comment:名称"`
	Data string `json:"data"  gorm:"column:data;type:longtext;comment:数值"`
	Desc string `json:"desc"  gorm:"column:desc;type:longtext;comment:数值"`
}

func (model *MediaOption) TableName() string {
	return "media_option"
}

func (model *MediaOption) GetOne(name string) (*MediaOption, error) {
	retVal := &MediaOption{}
	err := g_db.Where("name = ?", name).First(&retVal).Error
	return retVal, err
}

func (model *MediaOption) PatchOne(name string, data interface{}) error {

	result := g_db.Model(&MediaOption{}).Where("name = ?", name).Update("data", data)
	if result.RowsAffected == 0 {
		return errors.New("no rows were updated")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (model *MediaOption) NoFoundCreate(name string, data string, desc string) error {

	val := MediaOption{
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
		return err
	}

	return nil
}
