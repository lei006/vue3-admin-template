package model

import "errors"

// 报告模板 结构体  ReportTemplate
type ReportTemplate struct {
	BASE_MODEL
	Title       string `json:"title" form:"title" gorm:"column:title;comment:;size:tinytext;"`                   //title
	Template_id string `json:"template_id" form:"template_id" gorm:"column:template_id;comment:;size:tinytext;"` //模板ID
	Parent_id   string `json:"parent_id" form:"parent_id" gorm:"column:parent_id;comment:;size:tinytext;"`       //父节点ID
	Data_a      string `json:"data_a" form:"data_a" gorm:"column:data_a;comment:;size:longtext;type:text;"`      //data_a
	Data_b      string `json:"data_b" form:"data_b" gorm:"column:data_b;comment:;size:longtext;type:text;"`      //data_b
	Data_c      string `json:"data_c" form:"data_c" gorm:"column:data_c;comment:;size:longtext;type:text;"`      //data_c
	Data_d      string `json:"data_d" form:"data_d" gorm:"column:data_d;comment:;size:longtext;type:text;"`      //data_d
	Sort        string `json:"sort" form:"sort" gorm:"column:sort;comment:;size:tinytext;"`                      //排序
	AllowDelete *bool  `json:"allow_delete" form:"allow_delete" gorm:"column:allow_delete;comment:;"`            //允许删除
}

// TableName 报告模板 ReportTemplate自定义表名 report_template
func (ReportTemplate) TableName() string {
	return "report_template"
}

func (model *ReportTemplate) Create(template *ReportTemplate) error {
	err := g_db.Create(template).Error
	return err
}

func (model *ReportTemplate) DeleteOne(id uint) error {
	err := g_db.Delete(&SysUser{}, id).Error
	return err
}

func (model *ReportTemplate) DeleteMany(ids []uint) error {
	err := g_db.Delete(&[]ReportTemplate{}, "id in ?", ids).Error
	return err
}

func (model *ReportTemplate) UpdateOne(id uint, field string, data interface{}) error {

	result := g_db.Model(&ReportTemplate{}).Where("id = ?", id).Update(field, data)
	if result.RowsAffected == 0 {
		return errors.New("no rows were updated")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (model *ReportTemplate) GetOne(id uint) (*ReportTemplate, error) {

	val := &ReportTemplate{}
	err := g_db.Where("id = ?", id).First(val).Error

	return val, err
}

func (model *ReportTemplate) GetList() (list []ReportTemplate, total int64, err error) {

	// 创建db
	db := g_db.Model(&ReportTemplate{})
	var items []ReportTemplate
	err = db.Count(&total).Error
	if err != nil {
		return items, total, err
	}

	err = db.Find(&items).Error
	return items, total, err
}
