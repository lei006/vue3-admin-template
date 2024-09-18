package sign_tool

import (
	"testing"
	"vue3-admin-template/pkg/utils"

	"github.com/lei006/zlog"
)

func TestNewBase64Key(t *testing.T) {

	for val := range 100 {
		pub, pri, err := NewBase64Key()
		if err != nil {
			t.Error(err)
			return
		} else {
			check_str := utils.RandomString(3200000, true, true, true)
			sign, err := Base64Sign(check_str, pri)
			if err != nil {
				t.Error(err)
				return
			}

			ret, err := Base64Verify(check_str, sign, pub)
			if err != nil {
				t.Error(err)
				return
			}
			if ret != true {
				zlog.Error("Base64Verify error")
				return
			}

			zlog.Debug("pubkey:", pub)
			zlog.Debug("prikey:", pri)

			zlog.Debug(val, ": Base64Verify:", ret, " sign len=", len(sign))
		}

	}
}
