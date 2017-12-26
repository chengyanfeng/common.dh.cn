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

	Rpc = "https://rpc.dh.cn:8080"
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

func NotifyHandle(p P, auth string) {
	go func() {
		p["auth"] = auth
		req := httplib.Post(Rpc)
		req.Header("Content-Type", "application/json")
		req.Body([]byte(JsonEncode(p)))
		req.Response()
	}()
}
