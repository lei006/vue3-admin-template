package model

import (
	"errors"
)

type SysAdmin struct {
	BASE_MODEL
	Username string `json:"username" gorm:"uniqueIndex;comment:用户登录名"`      // 用户登录名
	Password string `json:"-"  gorm:"comment:用户登录密码"`                       // 用户登录密码
	Nickname string `json:"nickname" gorm:"default:系统用户;comment:用户昵称"`      // 用户昵称
	Token    string `json:"token" gorm:"index;default:token;comment:token"` // token
	//Headerimg string `json:"headerimg" gorm:"type:mediumtext; comment:用户头像"` // 用户头像
	UserSign string `json:"user_sign" gorm:"type:mediumtext;comment:用户签名"` // 用户签名
	Desc     string `json:"desc"  gorm:"comment:描述"`                       //
	//Email     string `json:"email"  gorm:"comment:用户邮箱"`                          // 用户邮箱
	IsDisable bool `json:"is_disable" gorm:"default:0;comment:用户是否被冻结 0正常 1冻结"` //用户是否被冻结 0正常 1冻结
}

func (SysAdmin) TableName() string {
	return "sys_admin"
}

// CreateReportStruct 创建报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysAdmin) Create(newVal *SysAdmin) (err error) {
	err = g_db.Create(newVal).Error
	return err
}

// DeleteReportStruct 删除报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysAdmin) DeleteOne(id string) (err error) {
	err = g_db.Delete(&SysAdmin{}, id).Error
	return err
}

// DeleteReportStructByIds 批量删除报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysAdmin) DeleteMany(ids []uint) (err error) {
	err = g_db.Delete(&[]SysAdmin{}, "id in ?", ids).Error
	return err
}

func (model *SysAdmin) UpdateOne(val *SysAdmin) (err error) {
	err = g_db.Model(val).Select("nickname", "phone", "user_sign").Save(val).Error
	return err
}

func (model *SysAdmin) PatchOne(id string, field string, data interface{}) error {

	result := g_db.Model(&SysAdmin{}).Where("id = ?", id).Update(field, data)
	if result.RowsAffected == 0 {
		return errors.New("no rows were updated")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (model *SysAdmin) GetField(id string, field string) (*SysAdmin, error) {
	//err = g_db.Where("id = ?", id).First(&retVal).Error
	retVal := &SysAdmin{}
	err := g_db.Where("id = ?", id).Select(field).First(retVal).Error
	return retVal, err
}

func (model *SysAdmin) GetOne(id string) (retVal SysAdmin, err error) {
	err = g_db.Where("id = ?", id).First(&retVal).Error
	return
}

func (model *SysAdmin) GetOneByUsername(username string) (retVal *SysAdmin, err error) {
	// 通过username 取得一行
	retVal = &SysAdmin{}
	result := g_db.Where("username = ?", username).First(retVal)
	if result.RowsAffected == 0 {
		return nil, errors.New("no rows were returned")
	}

	return retVal, result.Error
}

func (model *SysAdmin) GetOneByToken(token string) (retVal *SysAdmin, err error) {
	// 通过username 取得一行
	retVal = &SysAdmin{}
	err = g_db.Where("token = ?", token).First(retVal).Error
	return
}

// GetReportStructInfoList 分页获取报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysAdmin) GetPage() (list []SysAdmin, total int64, err error) {

	// 创建db
	db := g_db.Model(&SysAdmin{})
	var reportItems []SysAdmin
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&reportItems).Error
	return reportItems, total, err
}

func (model *SysAdmin) FindOrCreate(username string, password string) (*SysAdmin, error) {

	val := SysAdmin{
		Username: username,
		Password: password,
		Nickname: username,
	}

	result := g_db.Where("name = ?", username).First(&val)
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
