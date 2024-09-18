package db_model

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type LicenseStruct struct {
	AppName        string    `json:"appname" gorm:"index;comment:用户登录名"`                 // 用户登录名
	Company        string    `json:"company" gorm:"index;comment:公司"`                    // 用户登录名
	HardSn         string    `json:"hard_sn" gorm:"index;column:hard_sn;comment:硬件唯一标识"` //
	LicenseText0   string    `json:"license_text0" gorm:"type:mediumtext;comment:文本1"`
	LicenseText1   string    `json:"license_text1" gorm:"type:mediumtext;comment:文本2"`
	LicenseText2   string    `json:"license_text2" gorm:"type:mediumtext;comment:文本3"`
	LicenseLimit0  int       `json:"license_limit0" gorm:"comment:授权数1"`
	LicenseLimit1  int       `json:"license_limit1" gorm:"comment:授权数2"`
	LicenseLimit2  int       `json:"license_limit2" gorm:"comment:授权数3"`
	LicenseAt      time.Time `json:"license_at" gorm:"column:license_at;comment:授权时间"`                //
	LicenseTimeLen int       `json:"license_time_len" gorm:"column:license_time_len;comment:授权长度(月)"` //
	Desc           string    `json:"desc" gorm:"type:mediumtext;comment:授权时间"`                        //
	PubKey         string    `json:"pub_key" gorm:"type:mediumtext;"`
	Sign           string    `json:"sign" gorm:"column:sign;type:mediumtext;comment:签名"`
}

type SysLicense struct {
	BASE_MODEL
	LicenseStruct
	LicenseData string `json:"license_data" gorm:"column:license_data;type:longtext;comment:用户签名"`
	IsDisable   bool   `json:"is_disable" gorm:"column:is_disable;default:false;comment:是否被冻结 0正常 1冻结"` //是否被冻结 0正常 1冻结
}

func (SysLicense) TableName() string {
	return "license_data"
}

// CreateReportStruct 创建报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysLicense) Create(newVal *SysLicense) (err error) {
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
func (model *SysLicense) DeleteOne(id string) (err error) {
	err = g_db.Unscoped().Delete(&SysLicense{}, id).Error
	return err
}

// DeleteReportStructByIds 批量删除报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysLicense) DeleteMany(ids []int) (err error) {
	err = g_db.Unscoped().Delete(&[]SysLicense{}, "id in ?", ids).Error
	return err
}

func (model *SysLicense) UpdateOne(val *SysLicense) (err error) {
	err = g_db.Model(val).Select("username", "nickname", "password", "usersign", "desc", "is_disable").Save(val).Error
	return err
}

func (model *SysLicense) PatchOne(id string, field string, data interface{}) error {

	result := g_db.Model(&SysLicense{}).Where("id = ?", id).Update(field, data)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no field were updated")
	}

	return nil
}

func (model *SysLicense) GetField(id string, field string) (*SysLicense, error) {
	//err = g_db.Where("id = ?", id).First(&retVal).Error
	retVal := &SysLicense{}
	err := g_db.Where("id = ?", id).Select(field).First(retVal).Error
	return retVal, err
}

func (model *SysLicense) GetOne(id string) (retVal SysLicense, err error) {
	err = g_db.Where("id = ?", id).First(&retVal).Error
	return
}

func (model *SysLicense) GetOneByUsername(username string) (retVal *SysLicense, err error) {
	// 通过username 取得一行
	retVal = &SysLicense{}
	result := g_db.Where("username = ?", username).First(retVal)
	if result.RowsAffected == 0 {
		return nil, errors.New("no rows were returned")
	}

	return retVal, result.Error
}

func (model *SysLicense) GetOneByToken(token string) (retVal *SysLicense, err error) {
	// 通过username 取得一行
	retVal = &SysLicense{}
	err = g_db.Where("token = ?", token).First(retVal).Error
	return
}

// GetReportStructInfoList 分页获取报告的数据结构记录
// Author [piexlmax](https://github.com/piexlmax)
func (model *SysLicense) GetPage(page PageInfo) (list []SysLicense, total int64, err error) {

	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)

	// 创建db
	db := g_db.Model(&SysLicense{})

	if len(page.Keyword) > 0 {
		db = db.Where("username LIKE ?", "%"+page.Keyword+"%").Or("nickname LIKE ?", "%"+page.Keyword+"%").Or("desc LIKE ?", "%"+page.Keyword+"%")
	}

	var reportItems []SysLicense
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

func LicenseGetJson(model *LicenseStruct) (string, error) {
	data, err := json.Marshal(model)
	if err != nil {
		return "", err
	}
	return string(data), nil

}

func (model *LicenseStruct) GetData() string {

	str := model.AppName
	str += model.Company
	str += model.HardSn
	str += model.LicenseText0
	str += model.LicenseText1
	str += model.LicenseText2
	str += fmt.Sprintf("%d", model.LicenseLimit0)
	str += fmt.Sprintf("%d", model.LicenseLimit1)
	str += fmt.Sprintf("%d", model.LicenseLimit2)
	str += model.LicenseAt.Format(time.RFC3339)
	str += fmt.Sprintf("%d", model.LicenseTimeLen)

	return str
}
