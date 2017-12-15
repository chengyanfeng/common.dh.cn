package util

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/sha1"
	"sort"
	"encoding/json"
	"bytes"
)

var UC_CONFIG map[string]string = map[string]string{
	"public_key":  "ucloudsupport@mrocker.com1392263197892193080",
	"private_key": "cff5a64df861f90a91eba840c51bae8f44fe008b",
	"project_id":  "org-5875",
	"base_url":    "https://api.ucloud.cn",
}

type Ucloud struct {
}

func (this *Ucloud) SendSms(msg string, mobile ...string) (result string) {
	params := this.CommonParams()
	params["Action"] = "SendSms"
	params["Content"] = msg

	for key, val := range mobile {
		params["Phone."+string(key)] = val
	}
	params["Signature"] = this.VerfyAc(params, UC_CONFIG["private_key"])

	data, err := this.Request(UC_CONFIG["base_url"], params)
	if err != nil {
		Error(err)
	}
	result = data
	Debug(result)
	return
}

func (this *Ucloud) RefreshCdn(url string) (result string) {
	params := this.CommonParams()
	params["Action"] = "RefreshUcdnDomainCache"
	params["Type"] = "dir"
	params["DomainId"] = "ucdn-d11yag"
	params["UrlList.0"] = url
	params["Signature"] = this.VerfyAc(params, UC_CONFIG["private_key"])

	data, err := this.Request(UC_CONFIG["base_url"], params)
	if err != nil {
		Error(err)
	}
	result = data
	Debug(result)
	return
}

func (uc *Ucloud) CommonParams() map[string]string {
	params := map[string]string{}
	params["PublicKey"] = UC_CONFIG["public_key"]
	params["ProjectId"] = UC_CONFIG["project_id"]
	return params
}

func (uc *Ucloud) VerfyAc(params map[string]string, private_key string) string {
	params_data := ""

	sorted_keys := make([]string, 0)
	for key, _ := range params {
		sorted_keys = append(sorted_keys, key)
	}

	sort.Strings(sorted_keys)

	for _, v := range sorted_keys {
		params_data += v
		params_data += params[v]
	}

	params_data += private_key

	return fmt.Sprintf("%x", sha1.Sum([]byte(params_data)))
}

func (uc *Ucloud) Request(base_url string, params map[string]string) (string, error) {
	client := &http.Client{}
	b, err := json.Marshal(params)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", base_url, bytes.NewBuffer([]byte(b)))

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
