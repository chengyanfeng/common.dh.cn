package def

import (
	"time"
)

var (
	EE         bool   = false
	MODE       string = ""
	MONGO_HOST string = ""
	UPTIME            = time.Now().UnixNano() / int64(time.Millisecond)
	Md5Salt    string = "Dh@)!^o5l3!%Op0"

	Jdbc_proxy_url    string = "http://jdbc.datahunter.cn/sql"
	Fitting_proxy_url string = "http://jdbc.datahunter.cn/fitting"
	Jdbc_xinghuan_url string = "http://xinghuan.jdbc.datahunter.cn/sql"

	PUNCTUATION []string = []string{".", ";", ",", "(", ")", "%"}
)
