package sign_tool

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"log"
	"math/big"
)

func NewBase64Key() (string, string, error) {

	//生成密钥对
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		return "", "", err
	}

	//x509编码
	eccPrivateKey, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return "", "", err
	}

	//保存公钥
	publicKey := privateKey.PublicKey

	//x509编码
	eccPublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return "", "", errors.New("MarshalPKIXPublicKey error" + err.Error())
	}

	return base64.StdEncoding.EncodeToString(eccPublicKey), base64.StdEncoding.EncodeToString(eccPrivateKey), nil
}

func privateKeyFromBase64(private_str string) (*ecdsa.PrivateKey, error) {
	//读取私钥

	eccPrivateKey, err := base64.StdEncoding.DecodeString(private_str)
	if err != nil {
		return nil, err
	}

	//x509解码
	privateKey, err := x509.ParseECPrivateKey(eccPrivateKey)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func publicKeyFromBase64(publish_str string) (*ecdsa.PublicKey, error) {
	//读取公钥
	//pem解密
	eccPpublicKey, err := base64.StdEncoding.DecodeString(publish_str)
	if err != nil {
		return nil, err
	}

	//x509解密
	publicInterface, err := x509.ParsePKIXPublicKey(eccPpublicKey)
	if err != nil {
		return nil, err
	}
	publicKey := publicInterface.(*ecdsa.PublicKey)
	return publicKey, nil
}

// 对消息的散列值生成数字签名
func Base64Sign(msg string, private_str string) (string, error) {

	// 捕获异常
	defer func() {
		if err := recover(); err != nil {
			log.Println("Base64Sign error:", err)
		}
	}()

	//取得私钥
	privateKey, err := privateKeyFromBase64(private_str)
	if err != nil {
		return "", err
	}

	//计算哈希值
	hash := sha256.New()
	//填入数据
	hash.Write([]byte(msg))
	bytes := hash.Sum(nil)
	//对哈希值生成数字签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, bytes)
	if err != nil {
		return "", err
	}
	rtext, err := r.MarshalText()
	if err != nil {
		return "", err
	}
	stext, err := s.MarshalText()
	if err != nil {
		return "", err
	}

	sign := &EccSign{
		Rtext: string(rtext),
		Stext: string(stext),
	}

	tmp_str, err := sign.ToJson()
	if err != nil {
		return "", err
	}

	return tmp_str, nil
}

// 验证数字签名
func Base64Verify(msg string, sign string, publish_key_str string) (bool, error) {

	defer func() {
		if err := recover(); err != nil {
			log.Println("Base64Verify error:", err)
		}
	}()

	//读取公钥
	publicKey, err := publicKeyFromBase64(publish_key_str)
	if err != nil {
		return false, err
	}

	ecc_sign := EccSign{}
	err = ecc_sign.FromJson(sign)
	if err != nil {
		return false, err
	}
	//计算哈希值
	hash := sha256.New()
	hash.Write([]byte(msg))
	bytes := hash.Sum(nil)

	//验证数字签名
	var r, s big.Int
	r.UnmarshalText([]byte(ecc_sign.Rtext))
	s.UnmarshalText([]byte(ecc_sign.Stext))

	verify := ecdsa.Verify(publicKey, bytes, &r, &s)
	return verify, nil
}
