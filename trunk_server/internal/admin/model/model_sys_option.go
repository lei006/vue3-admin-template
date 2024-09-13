// 自动生成模板SysOperationRecord
package model

// 如果含有time.Time 请自行import time包
type SysOption struct {
	BASE_MODEL
	UserID    int    `json:"user_id" form:"user_id" gorm:"column:user_id;comment:用户id"`       // 用户id
	FromIp    string `json:"fromip" form:"fromip" gorm:"column:fromip;comment:请求ip"`          // 请求ip
	MsgText01 string `json:"msgtext01" form:"msgtext01" gorm:"column:msgtext01;comment:请求方法"` // 0
	MsgText02 string `json:"msgtext02" form:"msgtext02" gorm:"column:msgtext02;comment:请求方法"` // 0
	MsgText03 string `json:"msgtext03" form:"msgtext03" gorm:"column:msgtext03;comment:请求方法"` // 0
	MsgText04 string `json:"msgtext04" form:"msgtext04" gorm:"column:msgtext04;comment:请求方法"` // 0
	MsgText05 string `json:"msgtext05" form:"msgtext05" gorm:"column:msgtext05;comment:请求方法"` // 0
	MsgText06 string `json:"msgtext06" form:"msgtext06" gorm:"column:msgtext06;comment:请求方法"` // 0
	MsgText07 string `json:"msgtext07" form:"msgtext07" gorm:"column:msgtext07;comment:请求方法"` // 0
	MsgText08 string `json:"msgtext08" form:"msgtext08" gorm:"column:msgtext08;comment:请求方法"` // 0
	MsgText09 string `json:"msgtext09" form:"msgtext09" gorm:"column:msgtext09;comment:请求方法"` // 0
	MsgText10 string `json:"msgtext10" form:"msgtext10" gorm:"column:msgtext10;comment:请求方法"` // 0
}

func (SysOption) TableName() string {
	return "sys_option"
}

func (model *SysOption) DeleteMany(ids []int) (err error) {
	result := g_db.Unscoped().Delete(&[]SysOption{}, "id in ?", ids)
	if result.Error != nil {
		return result.Error
	}

	return err
}

func (model *SysOption) GetPage(page PageInfo) (list []SysOption, total int64, err error) {

	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)

	// 创建db
	db := g_db.Model(&SysOption{})

	if len(page.Keyword) > 0 {
		db = db.Where("username LIKE ?", "%"+page.Keyword+"%")
	}

	var items []SysOption
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
