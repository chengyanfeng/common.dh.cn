package auth

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	. "common.dh.cn/utils"
)

func Init() {
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
