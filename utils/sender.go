package utils

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

var (
	Monitor = "http://alert.datahunter.cn:16880/transfer/portal"
	Tos     = "WangChengLong"
)

func send(tos, appid, title, content, tp string, eventtime int64, merge int8) {
	if len(tos) == 0 {
		tos = Tos
	}

	monitor := Monitor
	str := fmt.Sprintf("%s?tos=%s&appid=%s&title=%s&content=%s&type=%s&eventtime=%d&merge=%d", monitor, tos, appid, title, url.QueryEscape(content), tp, eventtime, merge)

	resp, err := httplib.Get(str).Response()
	if err != nil {
		logs.Error("%s send: %v", err.Error(), str)

		return
	}

	Debug("==send==>>>>%v<<<<==send==%s", str, resp)
}

func Send(tos, title, content string) {
	now := time.Now().Unix()
	appid := GetHostname()

	send(tos, strings.TrimSpace(appid), title, content, "Error", now, 1)
}