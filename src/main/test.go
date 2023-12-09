package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// 生成数据库密钥
const PasswdEncodeKey = "kenc-wayne key 1"

func DecodePassword(encodeBase64 string) (string, error) {
	var pwd string
	key := []byte(PasswdEncodeKey)
	encode, err := base64.StdEncoding.DecodeString(encodeBase64)
	if err != nil {
		//logs.Error("parse DBPasswd error.%v", err)
		return pwd, err
	}

	pwdBytes, err := decryptAES(encode, key)
	if err != nil {
		//logs.Error("decrypt error.%v", err)
	}
	pwd = string(pwdBytes)
	return pwd, nil
}

func decryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		//logs.Error("aes cipher error.%v", err)
		return nil, err
	}
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return src, nil
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func main() {
	pwd, err := DecodePassword("QX2GihLM4A7dKK7YGdSCd39hry3OXOTNDcoUZhlcLpc=")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pwd)
}
