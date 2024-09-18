// 自动生成模板SysOperationRecord
package db_model

// 如果含有time.Time 请自行import time包
type SysAbout struct {
	BASE_MODEL
	Title string `json:"title" form:"title" gorm:"column:title;comment:"` // 0
	Data  string `json:"data" form:"data" gorm:"column:data;comment:"`    // 0
	Desc  string `json:"desc" form:"desc" gorm:"column:desc;comment:"`    // 0
}

func (SysAbout) TableName() string {
	return "sys_about"
}

func (model *SysAbout) GetPage(page PageInfo) (list []SysAbout, total int64, err error) {

	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)

	// 创建db
	db := g_db.Model(&SysAbout{})

	var items []SysAbout
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	result := db.Limit(limit).Offset(offset).Order("id desc").Find(&items)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return items, total, err
}

func FindOrCreateAbout(title string, data string, desc string) (*SysAbout, error) {

	val := SysAbout{
		Title: title,
		Data:  data,
		Desc:  desc,
	}
	result := g_db.Where("title = ?", title).First(&val)
	if result.RowsAffected == 0 {
		create_result := g_db.Create(&val)
		if create_result.Error != nil {
			return nil, create_result.Error
		}
		return &val, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &val, nil
}
