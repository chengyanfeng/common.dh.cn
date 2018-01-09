package utils

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/astaxie/beego/config"
)

var UCLOUND_CONFIG map[string]string

func init() {
	cnf, err := config.NewConfig("ini", "conf/ucloud.conf")
	if err != nil {
		Debug("ucloud配置文件失败")
	} else {
		UCLOUND_CONFIG, err = cnf.GetSection("ucloud")
		if err != nil {
			Debug("获取ucloud配置文件失败")
		}
	}
}

type Ucloud struct {
}

func (m *Ucloud) SendSms(msg string, mobile ...string) (result string) {
	params := P{}
	params["PublicKey"] = UCLOUND_CONFIG["public_key"]
	params["ProjectId"] = UCLOUND_CONFIG["project_id"]
	params["Action"] = "SendSms"
	params["Content"] = msg
	for key, val := range mobile {
		params["Phone."+string(key)] = val
	}
	params["Signature"] = m.VerfyAc(params, UCLOUND_CONFIG["private_key"])
	body, _ := json.Marshal(params)
	data, err := HttpPostBody(UCLOUND_CONFIG["base_url"], &P{}, body)
	if err != nil {
		Error("Ucloud Send Sms Error :", err)
	}
	return data
}

func (m *Ucloud) RefreshCdn(url string) (result string) {
	params := P{}
	params["PublicKey"] = UCLOUND_CONFIG["public_key"]
	params["ProjectId"] = UCLOUND_CONFIG["project_id"]
	params["Action"] = "RefreshUcdnDomainCache"
	params["Type"] = "dir"
	params["DomainId"] = "ucdn-d11yag"
	params["UrlList.0"] = url
	params["Signature"] = m.VerfyAc(params, UCLOUND_CONFIG["private_key"])
	body, _ := json.Marshal(params)
	data, err := HttpPostBody(UCLOUND_CONFIG["base_url"], &P{}, body)
	if err != nil {
		Error("Ucloud Refresh Cdn Error :", err)
	}
	return data
}

func (m *Ucloud) VerfyAc(params P, private_key string) string {
	params_data := ""
	sorted_keys := make([]string, 0)
	for key, _ := range params {
		sorted_keys = append(sorted_keys, key)
	}
	sort.Strings(sorted_keys)
	for _, v := range sorted_keys {
		params_data += v
		params_data += params[v].(string)
	}
	params_data += private_key
	return fmt.Sprintf("%x", sha1.Sum([]byte(params_data)))
}
