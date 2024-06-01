package model

import (
	"errors"

	"github.com/gofrs/uuid/v5"
)

type SysUser struct {
	BASE_MODEL
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`               // 用户UUID
	Username string    `json:"userName" gorm:"index;comment:用户登录名"`            // 用户登录名
	Password string    `json:"-"  gorm:"comment:用户登录密码"`                       // 用户登录密码
	NickName string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`      // 用户昵称
	Token    string    `json:"token" gorm:"index;default:token;comment:token"` // token
	SideMode string    `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`    // 用户侧边主题
	//HeaderImg   string         `json:"headerImg" gorm:"type:TEXT; default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	HeaderImg   string `json:"headerImg" gorm:"type:TEXT; comment:用户头像"`           // 用户头像
	UserSign    string `json:"userSign" gorm:"type:TEXT;comment:用户签名"`             // 用户签名
	BaseColor   string `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`         // 基础颜色
	ActiveColor string `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`    // 活跃颜色
	Authority   string `json:"authority" gorm:"default:888;comment:用户角色ID"`        // 用户角色ID
	Phone       string `json:"phone"  gorm:"comment:用户手机号"`                        // 用户手机号
	Email       string `json:"email"  gorm:"comment:用户邮箱"`                         // 用户邮箱
	Disenable   bool   `json:"disenable" gorm:"default:0;comment:用户是否被冻结 0正常 1冻结"` //用户是否被冻结 0正常 1冻结
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

func (model *SysUser) PatchOne(id uint, data interface{}) error {

	result := g_db.Model(&SysUser{}).Where("id = ?", id).Update("token", data)
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
