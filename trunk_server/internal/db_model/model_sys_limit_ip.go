package db_model

import (
	"errors"
)

// 配置文件结构体
type SysLimitIp struct {
	BASE_MODEL
	Ip      string `json:"ip" gorm:"column:ip;uniqueIndex;comment:名称"`
	IsLimit bool   `json:"is_limit" gorm:"column:is_limit;default:false;comment:是否限制 0不限 1限制"` //iP是否被限制 0正常 1冻结
	Desc    string `json:"desc" gorm:"column:desc;type:tinytext;comment:描述"`
}

func (model *SysLimitIp) TableName() string {
	return "sys_limit_ip"
}

func (model *SysLimitIp) Create(val *SysLimitIp) (*SysLimitIp, error) {
	result := g_db.Create(val)
	if result.RowsAffected == 0 {

		return nil, errors.New("no rows were create")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return val, result.Error
}

func (model *SysLimitIp) PatchOne(id string, field string, data interface{}) error {

	result := g_db.Model(&SysLimitIp{}).Where("id = ?", id).Update(field, data)
	if result.RowsAffected == 0 {
		return errors.New("no rows were updated")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (model *SysLimitIp) DeleteMany(ids []int) (err error) {
	result := g_db.Unscoped().Delete(&[]SysLimitIp{}, "id in ?", ids)
	if result.RowsAffected == 0 {
		return errors.New("no rows were delete")
	}

	if result.Error != nil {
		return result.Error
	}

	return err
}

func (model *SysLimitIp) IsExist(ip string) (bool, error) {
	// 通过username 取得一行
	retVal := &SysLimitIp{}
	result := g_db.Where("ip = ?", ip).First(retVal)
	// 判断是否存在
	if result.RowsAffected == 0 {
		return false, errors.New("ip is not exist:" + ip)
	}
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func (model *SysLimitIp) GetOne(id string) (*SysLimitIp, error) {

	retVal := SysLimitIp{}
	result := g_db.Where("id = ?", id).First(&retVal)
	// 判断是否存在
	if result.RowsAffected == 0 {
		return nil, errors.New("id is not get:" + id)
	}

	if result.Error != nil {
		return nil, result.Error
	}
	return &retVal, nil
}

func (model *SysLimitIp) GetOneIp(ip string) (*SysLimitIp, error) {

	retVal := SysLimitIp{}
	result := g_db.Where("ip = ?", ip).First(&retVal)

	if result.Error != nil {
		return nil, result.Error
	}
	return &retVal, nil
}

func (model *SysLimitIp) GetPage(page PageInfo) (items []SysLimitIp, total int64, err error) {

	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)

	// 创建db
	db := g_db.Model(&SysLimitIp{})

	if len(page.Keyword) > 0 {
		db = db.Where("ip LIKE ?", "%"+page.Keyword+"%").Or("desc LIKE ?", "%"+page.Keyword+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	result := db.Limit(limit).Offset(offset).Order("id desc").Find(&items)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return items, total, err
}
