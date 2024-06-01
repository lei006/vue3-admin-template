package model

import (
	"errors"
)

type SysUser struct {
	BASE_MODEL
	Username  string `json:"username" gorm:"index;comment:用户登录名"`                // 用户登录名
	Password  string `json:"-"  gorm:"comment:用户登录密码"`                           // 用户登录密码
	Nickname  string `json:"nickname" gorm:"default:系统用户;comment:用户昵称"`          // 用户昵称
	Token     string `json:"token" gorm:"index;default:token;comment:token"`     // token
	Headerimg string `json:"headerimg" gorm:"type:TEXT; comment:用户头像"`           // 用户头像
	UserSign  string `json:"user_sign" gorm:"type:TEXT;comment:用户签名"`            // 用户签名
	IsAdmin   bool   `json:"is_admin" gorm:"default:0;comment:是否管理员"`            // 用户角色ID
	Phone     string `json:"phone"  gorm:"comment:用户手机号"`                        // 用户手机号
	Email     string `json:"email"  gorm:"comment:用户邮箱"`                         // 用户邮箱
	Disenable bool   `json:"disenable" gorm:"default:0;comment:用户是否被冻结 0正常 1冻结"` //用户是否被冻结 0正常 1冻结
}

func (SysUser) TableName() string {
	return "sys_user"
}

// CreateReportStruct 创建报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysUser) Create(newVal *SysUser) (err error) {
	err = g_db.Create(newVal).Error
	return err
}

// DeleteReportStruct 删除报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysUser) DeleteOne(id string) (err error) {
	err = g_db.Delete(&SysUser{}, id).Error
	return err
}

// DeleteReportStructByIds 批量删除报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysUser) DeleteMany(ids []uint) (err error) {
	err = g_db.Delete(&[]SysUser{}, "id in ?", ids).Error
	return err
}

// UpdateReportStruct 更新报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysUser) UpdateOne(val *SysUser) (err error) {
	err = g_db.Save(&val).Error
	return err
}

func (model *SysUser) PatchOne(id string, field string, data interface{}) error {

	result := g_db.Model(&SysUser{}).Where("id = ?", id).Update(field, data)
	if result.RowsAffected == 0 {
		return errors.New("No rows were updated")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetReportStruct 根据id获取报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysUser) GetOne(id string) (retVal SysUser, err error) {
	err = g_db.Where("id = ?", id).First(&retVal).Error
	return
}

func (model *SysUser) GetOneByUsername(username string) (retVal *SysUser, err error) {
	// 通过username 取得一行
	retVal = &SysUser{}
	err = g_db.Where("username = ?", username).First(retVal).Error
	return
}

func (model *SysUser) GetOneByToken(token string) (retVal *SysUser, err error) {
	// 通过username 取得一行
	retVal = &SysUser{}
	err = g_db.Where("token = ?", token).First(retVal).Error
	return
}

// GetReportStructInfoList 分页获取报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysUser) GetList() (list []SysUser, total int64, err error) {

	// 创建db
	db := g_db.Model(&SysUser{})
	var reportItems []SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&reportItems).Error
	return reportItems, total, err
}
