package license

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"vue3-admin-template/internal/config"
	"vue3-admin-template/internal/db_model"
	"vue3-admin-template/pkg/sign_tool"
	"vue3-admin-template/pkg/utils"

	"github.com/lei006/zlog"
)

func Init() error {

	///////////////////////////////////////////////////////
	// 取得硬件ID
	hardsn, err := utils.HardSn()
	if err != nil {
		return fmt.Errorf("获取硬件ID失败:%s", err.Error())
	}
	config.HardSn = hardsn
	zlog.Info("HardSn", hardsn)

	licenseFilepath := fmt.Sprintf("%s/%s", config.WorkPath, config.LicenseFile)

	// 验证文件是否存在
	if _, err := os.Stat(licenseFilepath); os.IsNotExist(err) {
		//不存在，则直接返回
		zlog.Warn("license.lic文件不存在，不进行授权验证", licenseFilepath)
		return nil
	}

	// 读取文件 licenseFilepath的内容
	license_file_data, err := os.ReadFile(licenseFilepath)
	if err != nil {
		return fmt.Errorf("ReadFile失败: %s", err.Error())
	}
	license_json := string(license_file_data)

	bret, err := VerifyAndUpdate(license_json)
	if err != nil {
		return fmt.Errorf("LicenseVerify失败: %s", err.Error())
	}
	if !bret {
		zlog.Warn(fmt.Errorf("授权验证失败"))
		return nil
	}

	return nil
}

func SaveLicenseFile(data string) error {
	licenseFilepath := fmt.Sprintf("%s/%s", config.WorkPath, config.LicenseFile)

	err := os.WriteFile(licenseFilepath, []byte(data), os.FileMode(0644))
	if err != nil {
		zlog.Error(err)
		return err
	}
	return nil
}

func Sign(model *db_model.LicenseStruct, pri string) (sign string, err error) {

	str := model.GetData()
	//str += model.Desc

	sign_str, err := sign_tool.Base64Sign(str, pri)
	if err != nil {
		return "", err
	}

	return sign_str, err
}

func LoadLicenseData(license_json string) error {

	return nil
}

func verify(license_json string) (bool, *db_model.LicenseStruct, error) {

	var license db_model.LicenseStruct
	err := json.Unmarshal([]byte(license_json), &license)
	if err != nil {
		return false, nil, errors.New("解码出错：" + err.Error())
	}

	license_data := license.GetData()
	bret, err := sign_tool.Base64Verify(license_data, license.Sign, license.PubKey)
	if err != nil {
		return false, nil, errors.New("签名验证出错" + err.Error())
	}

	if !bret {
		return false, nil, nil
	}

	return bret, &license, nil
}

func VerifyOnly(license_json string) (bool, error) {

	bret, _, err := verify(license_json)
	if err != nil {
		return false, err
	}

	return bret, nil
}

func VerifyAndUpdate(license_json string) (bool, error) {

	bret, lic, err := verify(license_json)
	if err != nil {
		return false, err
	}
	if lic.AppName != config.AppName {
		return false, errors.New("授权名与本应用不匹配")
	}

	if lic.HardSn != config.HardSn {
		zlog.Error("硬件ID", config.HardSn)
		zlog.Error("授权ID", lic.HardSn)
		return false, errors.New("硬件ID与本机器不匹配")
	}

	// 验证通过，记录license
	config.LicenseCheck = true
	config.Lic = *lic

	return bret, nil
}
