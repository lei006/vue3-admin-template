package db_model

import (
	"errors"
)

type ModelPriceStruct struct {
	BASE_MODEL
	ProjectId uint   `json:"project_id" gorm:"column:project_id;comment:项目ID"`                        // 项目ID  project_id
	Title     string `json:"title" gorm:"column:title;;comment:标题"`                                   // 项目ID  project_id
	Price     int    `json:"price" gorm:"comment:价格(元)"`                                              //
	IsDisable bool   `json:"is_disable" gorm:"column:is_disable;default:false;comment:是否被冻结 0正常 1冻结"` //是否被冻结 0正常 1冻结
	Remark    string `json:"remark" gorm:"comment:备注"`
}

func (ModelPriceStruct) TableName() string {
	return "pro_price"
}

// CreateReportStruct 创建报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *ModelPriceStruct) Create(newVal *ModelPriceStruct) (err error) {
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
func (model *ModelPriceStruct) DeleteMany(ids []int) (err error) {
	err = g_db.Unscoped().Delete(&[]ModelPriceStruct{}, "id in ?", ids).Error
	return err
}

func (model *ModelPriceStruct) PatchOne(id string, field string, data interface{}) error {

	result := g_db.Model(&ModelPriceStruct{}).Where("id = ?", id).Update(field, data)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no field were updated")
	}

	return nil
}

func (model *ModelPriceStruct) GetField(id string, field string) (*ModelPriceStruct, error) {
	//err = g_db.Where("id = ?", id).First(&retVal).Error
	retVal := &ModelPriceStruct{}
	err := g_db.Where("id = ?", id).Select(field).First(retVal).Error
	return retVal, err
}

func (model *ModelPriceStruct) GetOne(id string) (retVal ModelPriceStruct, err error) {
	err = g_db.Where("id = ?", id).First(&retVal).Error
	return
}

// GetReportStructInfoList 分页获取报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *ModelPriceStruct) GetPage(project_id string) (list []ModelPriceStruct, total int64, err error) {

	// 创建db
	db := g_db.Model(&ModelPriceStruct{})

	if project_id != "" {
		db = db.Where("project_id = ?", project_id)
	}

	var reportItems []ModelPriceStruct
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	result := db.Order("title desc").Find(&reportItems)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return reportItems, total, err
}
