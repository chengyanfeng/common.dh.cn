package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"common.dh.cn/def"
	"github.com/astaxie/beego"
	"github.com/henrylee2cn/mahonia"
)

var SecretPath string
var LicensePath string

type License struct {
	Address string
	Number  int
	Expire  time.Time
}

type Secret struct {
	Number int
	Days   int
	Key    string
}

func GetHostname() (host string) {
	host, _ = os.Hostname()
	if IsEmpty(host) {
		host = "www.datahunter.cn"
	}
	Debug("host", host)
	return host
}

func init() {
	SecretPath = beego.AppConfig.DefaultString("Secret", "secret.key")
	LicensePath = beego.AppConfig.DefaultString("License", "license.key")
}

func GetAuthNumber() int {
	license, err := GetLicense()
	if err == nil {
		return license.Number
	}
	return 10
}

func Etho() (str string) {
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		if inter.Name == "eth0" || inter.Name == "en0" {
			return inter.HardwareAddr.String()
		}
	}
	return ""
}

//GetSecret 获取秘钥信息
func GetSecret() (secret *Secret, err error) {
	if !FileExists(SecretPath) {
		return nil, errors.New("secret.key不存在")
	}
	data := ReadFile(SecretPath)
	if data == "" {
		return nil, errors.New("secret.key内容为空")
	}
	decrypt, _ := DesDecrypt(string(data), def.DESSalt)
	if decrypt == "" {
		return nil, errors.New("secret.key解析异常")
	}
	secretInfo := strings.Split(decrypt, "|")
	if len(secretInfo) != 3 {
		return nil, errors.New("secret.key解析异常")
	}
	_secret := Secret{}
	_secret.Number, err = strconv.Atoi(secretInfo[0])
	if err != nil {
		return nil, errors.New("secret.key解析异常")
	}
	_secret.Days, err = strconv.Atoi(secretInfo[1])
	if err != nil {
		return nil, errors.New("secret.key解析异常")
	}
	if len([]rune(secretInfo[2])) != 8 {
		return nil, errors.New("secret.key解析异常")
	}
	_secret.Key = secretInfo[2]
	return &_secret, nil
}

//GetLicense 获取授权信息
func GetLicense() (license *License, err error) {
	data := ReadFile(LicensePath)
	if data == "" {
		return nil, errors.New("license.key内容为空")
	}
	secret, err := GetSecret()
	if err != nil {
		return nil, err
	}
	decrypt, _ := DesDecrypt(string(data), []byte(secret.Key))
	if decrypt == "" {
		return nil, errors.New("license.key解析异常")
	}
	license = &License{}
	err = json.Unmarshal([]byte(decrypt), license)
	if err != nil {
		return nil, errors.New("license.key内容异常")
	}
	return license, nil
}

//GenerateLicense 生成授权信息
func GenerateLicense(address string) error {
	secret, err := GetSecret()
	if err != nil {
		return err
	}
	license := License{}
	license.Address = address
	license.Number = secret.Number
	license.Expire = time.Now().AddDate(0, 0, secret.Days)
	licenseInfo, err := json.Marshal(license)
	encrypt, err := DesEncrypt(licenseInfo, []byte(secret.Key))
	if err != nil {
		return errors.New("license.key生成失败")
	}
	err = WriteFile(LicensePath, []byte(encrypt))
	if err != nil {
		return errors.New("license.key生成失败")
	}
	return nil
}

//DesEncrypt DES加密
func DesEncrypt(origData []byte, key []byte) (string, error) {
	//UTF-8 to GBK
	var enc mahonia.Encoder
	enc = mahonia.NewEncoder("gbk")
	origDataStr := enc.ConvertString(string(origData))
	origData = []byte(origDataStr)
	block, err := des.NewCipher(key)
	if err != nil {
		Error(err.Error())
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

//DesDecrypt DES解密
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

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
