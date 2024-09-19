package db_model

import (
	"errors"
)

type OrderStruct struct {
}

type ModelOrderStruct struct {
	BASE_MODEL
	UserId        uint   `json:"user_id" gorm:"column:user_id;comment:项目ID"`        // 项目ID  project_id
	ProjectId     uint   `json:"project_id" gorm:"column:project_id;;comment:项目ID"` // 项目ID  project_id
	PriceId       uint   `json:"price_id" gorm:"column:price_id;comment:项目ID"`      // 价格ID  project_id
	AppName       string `json:"appname" gorm:"index;comment:用户登录名"`                // 用户登录名
	LicenseLimit0 int    `json:"license_limit0" gorm:"comment:授权数1"`
	LicenseData   string `json:"license_data" gorm:"column:license_data;type:longtext;comment:用户签名"`
	IsDisable     bool   `json:"is_disable" gorm:"column:is_disable;default:false;comment:是否被冻结 0正常 1冻结"` //是否被冻结 0正常 1冻结
}

func (ModelOrderStruct) TableName() string {
	return "pro_order"
}

// CreateReportStruct 创建报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *ModelOrderStruct) Create(newVal *ModelOrderStruct) (err error) {
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
func (model *ModelOrderStruct) DeleteMany(ids []int) (err error) {
	err = g_db.Unscoped().Delete(&[]ModelOrderStruct{}, "id in ?", ids).Error
	return err
}

func (model *ModelOrderStruct) PatchOne(id string, field string, data interface{}) error {

	result := g_db.Model(&ModelOrderStruct{}).Where("id = ?", id).Update(field, data)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no field were updated")
	}

	return nil
}

func (model *ModelOrderStruct) GetOne(id string) (retVal ModelOrderStruct, err error) {
	err = g_db.Where("id = ?", id).First(&retVal).Error
	return
}

// GetReportStructInfoList 分页获取报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *ModelOrderStruct) GetPage(page PageInfo) (list []ModelOrderStruct, total int64, err error) {

	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)

	// 创建db
	db := g_db.Model(&ModelOrderStruct{})

	if len(page.Keyword) > 0 {
		db = db.Where("username LIKE ?", "%"+page.Keyword+"%").Or("nickname LIKE ?", "%"+page.Keyword+"%").Or("desc LIKE ?", "%"+page.Keyword+"%")
	}

	var reportItems []ModelOrderStruct
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
