package model

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
///
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type ReportItem struct {
	BASE_MODEL
	Report_id string `json:"report_id" form:"report_id" gorm:"column:report_id;comment:报告ID;size:128;"`         //报告ID
	Innertext string `json:"innertext" form:"innertext" gorm:"column:innertext;size:longtext;comment:报告的文本内容;"` //文本内容
	InnerHtml string `json:"innerHtml" form:"innerHtml" gorm:"column:innerHtml;size:longtext;comment:报告内容;"`    //报告内容
}

// TableName 报告的数据结构 ReportStruct自定义表名 report_struct
func (ReportItem) TableName() string {
	return "report_struct"
}

var InsReportItem = new(ReportItem)

// CreateReportStruct 创建报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *ReportItem) Create(newVal *ReportItem) (err error) {
	err = g_db.Create(newVal).Error
	return err
}

// DeleteReportStruct 删除报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *ReportItem) DeleteOne(id string) (err error) {
	err = g_db.Delete(&ReportItem{}, id).Error
	return err
}

// DeleteReportStructByIds 批量删除报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *ReportItem) DeleteMany(ids []uint) (err error) {
	err = g_db.Delete(&[]ReportItem{}, "id in ?", ids).Error
	return err
}

// UpdateReportStruct 更新报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *ReportItem) UpdateOne(val *ReportItem) (err error) {
	err = g_db.Save(&val).Error
	return err
}

// GetReportStruct 根据id获取报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *ReportItem) GetOne(id string) (retVal ReportItem, err error) {
	err = g_db.Where("id = ?", id).First(&retVal).Error
	return
}

// GetReportStructInfoList 分页获取报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *ReportItem) GetList() (list []ReportItem, total int64, err error) {

	// 创建db
	db := g_db.Model(&ReportItem{})
	var reportItems []ReportItem
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&reportItems).Error
	return reportItems, total, err
}
