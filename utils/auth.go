package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
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
