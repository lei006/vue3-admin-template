package db_model

import (
	"errors"
)

type ModelProjectStruct struct {
	BASE_MODEL
	ProjectName string `json:"project_name" gorm:"column:project_name;comment:项目名称"` // 项目名称
	Remark      string `json:"remark" gorm:"comment:备注"`
	IsDisable   bool   `json:"is_disable" gorm:"column:is_disable;default:false;comment:是否被冻结 0正常 1冻结"` //是否被冻结 0正常 1冻结
}

func (ModelProjectStruct) TableName() string {
	return "pro_project"
}

// CreateReportStruct 创建报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *ModelProjectStruct) Create(newVal *ModelProjectStruct) (err error) {
	result := g_db.Create(newVal)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no rows were inserted")
	}

	return nil
}

// DeleteReportStructByIds 批量删除报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *ModelProjectStruct) DeleteMany(ids []int) (err error) {
	err = g_db.Unscoped().Delete(&[]ModelProjectStruct{}, "id in ?", ids).Error
	return err
}

func (model *ModelProjectStruct) PatchOne(id string, field string, data interface{}) error {

	result := g_db.Model(&ModelProjectStruct{}).Where("id = ?", id).Update(field, data)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no field were updated")
	}

	return nil
}

func (model *ModelProjectStruct) GetField(id string, field string) (*ModelProjectStruct, error) {
	//err = g_db.Where("id = ?", id).First(&retVal).Error
	retVal := &ModelProjectStruct{}
	err := g_db.Where("id = ?", id).Select(field).First(retVal).Error
	return retVal, err
}

func (model *ModelProjectStruct) GetOne(id string) (retVal ModelProjectStruct, err error) {
	err = g_db.Where("id = ?", id).First(&retVal).Error
	return
}

// GetReportStructInfoList 分页获取报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *ModelProjectStruct) GetPage(page PageInfo) (list []ModelProjectStruct, total int64, err error) {

	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)

	// 创建db
	db := g_db.Model(&ModelProjectStruct{})

	var reportItems []ModelProjectStruct
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	result := db.Limit(limit).Offset(offset).Order("id desc").Find(&reportItems)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return reportItems, total, err
}
