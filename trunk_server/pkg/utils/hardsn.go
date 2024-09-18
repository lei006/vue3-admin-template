package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/super-l/machine-code/machine"
)

func HardSn() (string, error) {

	machineData := machine.GetMachineData()
	result, err := json.Marshal(machineData)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	//对 machineData,做md5
	//zlog.Debug("machineData:", string(result))

	hasher := md5.New()
	hasher.Write([]byte(result))
	tmp_hard_sn := hex.EncodeToString(hasher.Sum(nil))
	return tmp_hard_sn[:15], nil
}
