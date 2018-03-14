package utils

import (
	"github.com/astaxie/beego"
)

// var (
// 	Monitor = "http://alert.datahunter.cn:16880/transfer/portal"
// 	Tos     = "WangChengLong"
// )

// func send(tos, appid, title, content, tp string, eventtime int64, merge int8) {
// 	if len(tos) == 0 {
// 		tos = Tos
// 	}
// 	monitor := Monitor
// 	str := fmt.Sprintf("%s?tos=%s&appid=%s&title=%s&content=%s&type=%s&eventtime=%d&merge=%d", monitor, tos, appid, title, url.QueryEscape(content), tp, eventtime, merge)
// 	resp, err := httplib.Get(str).Response()
// 	if err != nil {
// 		logs.Error("%s send: %v", err.Error(), str)

// 		return
// 	}
// }

// func Send(tos, title, content string) {
// 	now := time.Now().Unix()
// 	appid := GetHostname()

// 	send(tos, strings.TrimSpace(appid), title, content, "Error", now, 1)
// }

func NotifyHandle(uid string) {
	go func() {
		notify := beego.AppConfig.DefaultString("notify", "http://localhost:8001/v2/notify/receive")
		HttpPost(notify, nil, &P{"uid": uid})
	}()
}
