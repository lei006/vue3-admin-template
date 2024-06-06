package model

import "errors"

// 预置字段 结构体  PreStruct
type ReportPrefield struct {
	BASE_MODEL
	PreType   string `json:"field_type" form:"field_type" gorm:"column:field_type;comment:;"`          //字段类型
	FieldName string `json:"field_name" form:"field_name" gorm:"column:title;comment:;size:tinytext;"` //字段名称
	ShowName  string `json:"show_name" form:"show_name" gorm:"column:show_name;comment:tinytext;"`     //显示名称
	PreData   string `json:"pre_data" form:"pre_data" gorm:"column:pre_data;comment:;size:longtext;"`  //预置数据
	DefData   string `json:"def_data" form:"def_data" gorm:"column:def_data;comment:;size:tinydata;"`  //缺省数据
	Sort      string `json:"sort" form:"sort" gorm:"column:sort;comment:;size:tinytext;"`              //排序
	Desc      string `json:"desc" form:"desc" gorm:"column:desc;comment:;size:tinytext;"`              //描述
}

// TableName 预置字段 PreStruct自定义表名 pre_struct
func (ReportPrefield) TableName() string {
	return "report_prefield"
}

func (model *ReportPrefield) Create(template *ReportPrefield) error {
	err := g_db.Create(template).Error
	return err
}

func (model *ReportPrefield) DeleteOne(id uint) error {
	err := g_db.Delete(&SysUser{}, id).Error
	return err
}

func (model *ReportPrefield) DeleteMany(ids []uint) error {
	err := g_db.Delete(&[]ReportPrefield{}, "id in ?", ids).Error
	return err
}

func (model *ReportPrefield) UpdateOne(id uint, field string, data interface{}) error {

	result := g_db.Model(&ReportPrefield{}).Where("id = ?", id).Update(field, data)
	if result.RowsAffected == 0 {
		return errors.New("no rows were updated")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (model *ReportPrefield) GetOne(id uint) (*ReportPrefield, error) {

	val := &ReportPrefield{}
	err := g_db.Where("id = ?", id).First(val).Error

	return val, err
}

func (model *ReportPrefield) GetList() (list []ReportPrefield, total int64, err error) {

	// 创建db
	db := g_db.Model(&ReportPrefield{})
	var items []ReportPrefield
	err = db.Count(&total).Error
	if err != nil {
		return items, total, err
	}

	err = db.Find(&items).Error
	return items, total, err
}
