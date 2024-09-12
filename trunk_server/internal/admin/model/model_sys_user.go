package model

import (
	"errors"
)

type SysUser struct {
	BASE_MODEL
	Username string `json:"username" gorm:"uniqueIndex;comment:用户登录名"`      // 用户登录名
	Password string `json:"password"  gorm:"comment:用户登录密码"`                // 用户登录密码
	Nickname string `json:"nickname" gorm:"default:系统用户;comment:用户昵称"`      // 用户昵称
	Token    string `json:"token" gorm:"index;default:token;comment:token"` // token
	//Headerimg string `json:"headerimg" gorm:"type:mediumtext; comment:用户头像"` // 用户头像
	UserSign string `json:"usersign" gorm:"type:mediumtext;comment:用户签名"` // 用户签名
	Desc     string `json:"desc"  gorm:"comment:描述"`                      //
	//Email     string `json:"email"  gorm:"comment:用户邮箱"`                          // 用户邮箱
	IsDisable bool `json:"is_disable" gorm:"default:false;comment:用户是否被冻结 0正常 1冻结"` //用户是否被冻结 0正常 1冻结
}

func (SysUser) TableName() string {
	return "sys_user"
}

// CreateReportStruct 创建报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysUser) Create(newVal *SysUser) (err error) {
	result := g_db.Create(newVal)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no rows were inserted")
	}

	return nil
}

// DeleteReportStruct 删除报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysUser) DeleteOne(id string) (err error) {
	err = g_db.Delete(&SysUser{}, id).Error
	return err
}

// DeleteReportStructByIds 批量删除报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysUser) DeleteMany(ids []int) (err error) {
	err = g_db.Delete(&[]SysUser{}, "id in ?", ids).Error
	return err
}

func (model *SysUser) UpdateOne(val *SysUser) (err error) {
	err = g_db.Model(val).Select("username", "nickname", "password", "usersign", "desc", "is_disable").Save(val).Error
	return err
}

func (model *SysUser) PatchOne(id string, field string, data interface{}) error {

	result := g_db.Model(&SysUser{}).Where("id = ?", id).Update(field, data)
	if result.RowsAffected == 0 {
		return errors.New("no rows were updated")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (model *SysUser) GetField(id string, field string) (*SysUser, error) {
	//err = g_db.Where("id = ?", id).First(&retVal).Error
	retVal := &SysUser{}
	err := g_db.Where("id = ?", id).Select(field).First(retVal).Error
	return retVal, err
}

func (model *SysUser) GetOne(id string) (retVal SysUser, err error) {
	err = g_db.Where("id = ?", id).First(&retVal).Error
	return
}

func (model *SysUser) GetOneByUsername(username string) (retVal *SysUser, err error) {
	// 通过username 取得一行
	retVal = &SysUser{}
	result := g_db.Where("username = ?", username).First(retVal)
	if result.RowsAffected == 0 {
		return nil, errors.New("no rows were returned")
	}

	return retVal, result.Error
}

func (model *SysUser) GetOneByToken(token string) (retVal *SysUser, err error) {
	// 通过username 取得一行
	retVal = &SysUser{}
	err = g_db.Where("token = ?", token).First(retVal).Error
	return
}

// GetReportStructInfoList 分页获取报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysUser) GetPage(page PageInfo) (list []SysUser, total int64, err error) {

	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)

	// 创建db
	db := g_db.Model(&SysUser{})

	if len(page.Keyword) > 0 {
		db = db.Where("username LIKE ?", "%"+page.Keyword+"%")
	}

	var reportItems []SysUser
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
