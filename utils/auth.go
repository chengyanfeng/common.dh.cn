package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/henrylee2cn/mahonia"
)

func InitAuth() {
	var rand = "DhKey@)!&!@)!"
	net := md5.New()
	net.Write([]byte(Etho()))
	s := net.Sum(nil)
	str := hex.EncodeToString(s)
	_, err := os.Stat("license.key")
	if err == nil {
		p := map[string]interface{}{}
		data, _ := ioutil.ReadFile("license.key")
		json.Unmarshal([]byte(data), &p)
		key := fmt.Sprintf("%v", p["key"])
		number := fmt.Sprintf("%v", p["user"])
		if Sha_256(str+number+rand) != key {
			fmt.Println("AuthError")
			os.Exit(2)
		}
	} else {
		fmt.Println("AuthError")
		os.Exit(2)
	}
}

func GetHostname() (host string) {
	host, _ = os.Hostname()
	if IsEmpty(host) {
		host = "www.datahunter.cn"
	}
	Debug("host", host)
	return host
}

func GetAuthNumber() int {
	_, err := os.Stat("license.key")
	if err == nil {
		p := map[string]interface{}{}
		data, _ := ioutil.ReadFile("license.key")
		json.Unmarshal([]byte(data), &p)
		number := fmt.Sprintf("%v", p["user"])
		return ToInt(number)
	}
	return 10
}

func Etho() (str string) {
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		if inter.Name == "eth0" {
			str = inter.HardwareAddr.String()
		}
	}
	return
}

func Sha_256(sha_str string) (str string) {
	s := sha256.New()
	s.Write([]byte(sha_str))
	ss := s.Sum(nil)
	str = hex.EncodeToString(ss)
	return
}

//DES加密
func DesEncrypt(origData, key []byte) (string, error) {
	//UTF-8 to GBK
	var enc mahonia.Encoder
	enc = mahonia.NewEncoder("gbk")
	origDataStr := enc.ConvertString(string(origData))
	origData = []byte(origDataStr)

	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	//crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	//base64加密
	encodeString := base64.StdEncoding.EncodeToString(crypted)
	return encodeString, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//DES解密
func DesDecrypt(encodeString string, key []byte) (string, error) {
	var dec mahonia.Decoder
	//base64解密
	crypted, err := base64.StdEncoding.DecodeString(encodeString)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	//origData = PKCS5UnPadding(origData)

	origData = ZeroUnPadding(origData)
	//GBK to UTF-8
	dec = mahonia.NewDecoder("gbk")
	origDataStr := dec.ConvertString(string(origData))

	return origDataStr, nil
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
