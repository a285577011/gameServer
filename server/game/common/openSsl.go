package common

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"github.com/name5566/leaf/log"
)

type openSsl struct {
}

func NewOpenSsl() *openSsl {
	obj := &openSsl{}
	return obj
}
func DecryptOpenssl(str string, key string) map[string]interface{} {
	var ret map[string]interface{}
	str = DecryptOnlyOpenssl(str, key)
	if str == "" {
		return ret
	}
	//初始化一个空的字节数组,这里的16为AES-128-CBC的要求的key的长度
	keyByte := [16]byte{}
	//转换key为字节数组
	keyByteTemp := []byte(key)
	//依次赋值,这里的3为key的len
	for i := 0; i < len(keyByteTemp); i++ {
		keyByte[i] = keyByteTemp[i]
	}
	decByte := DecryptAes128Ecb([]byte(str), keyByte[:])
	json.Unmarshal(decByte, &ret)
	return ret
}
func DecryptAes128Ecb(data, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err!=nil{
		log.Debug("DecryptAes128Ecb %s",err)
		return []byte("")
	}
	decrypted := make([]byte, len(data))
	size := 16

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}
	decrypted=PKCS5UnPadding(decrypted)
	return decrypted
}
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
func DecryptOnlyOpenssl(str string, key string) string {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	deCodestr := string(decoded)
	return deCodestr
}
