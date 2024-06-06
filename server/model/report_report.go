package model

import "errors"

type ReportReportStruct struct {
	BASE_MODEL
	Report_id string `json:"report_id" form:"report_id" gorm:"column:report_id;comment:报告ID;size:128;"`         //报告ID
	Innertext string `json:"innertext" form:"innertext" gorm:"column:innertext;size:longtext;comment:报告的文本内容;"` //文本内容
	InnerHtml string `json:"innerHtml" form:"innerHtml" gorm:"column:innerHtml;size:longtext;comment:报告内容;"`    //报告内容
}

// TableName 报告的数据结构 ReportStruct自定义表名 report_struct
func (ReportReportStruct) TableName() string {
	return "report_report"
}

func (model *ReportReportStruct) Create(template *ReportReportStruct) error {
	err := g_db.Create(template).Error
	return err
}

func (model *ReportReportStruct) DeleteOne(id uint) error {
	err := g_db.Delete(&SysUser{}, id).Error
	return err
}

func (model *ReportReportStruct) DeleteMany(ids []uint) error {
	err := g_db.Delete(&[]ReportReportStruct{}, "id in ?", ids).Error
	return err
}

func (model *ReportReportStruct) UpdateOne(id uint, field string, data interface{}) error {

	result := g_db.Model(&ReportReportStruct{}).Where("id = ?", id).Update(field, data)
	if result.RowsAffected == 0 {
		return errors.New("no rows were updated")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (model *ReportReportStruct) GetOne(id uint) (*ReportReportStruct, error) {

	val := &ReportReportStruct{}
	err := g_db.Where("id = ?", id).First(val).Error

	return val, err
}

func (model *ReportReportStruct) GetPage() (list []ReportReportStruct, total int64, err error) {

	// 创建db
	db := g_db.Model(&ReportReportStruct{})
	var items []ReportReportStruct
	err = db.Count(&total).Error
	if err != nil {
		return items, total, err
	}

	err = db.Find(&items).Error
	return items, total, err
}
