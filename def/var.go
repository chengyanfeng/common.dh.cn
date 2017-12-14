package def

import (
	"time"

	"github.com/nanobox-io/golang-scribble"
	"github.com/orcaman/concurrent-map"
	"gopkg.in/robfig/cron.v2"

	. "common.dh.cn/util"
)

var (
	EE         bool   = false
	MODE       string = ""
	MONGO_HOST string = ""
	UPTIME            = time.Now().UnixNano() / int64(time.Millisecond)
	Md5Salt    string = "Dh@)!^o5l3!%Op0"

	Jdbc_proxy_url    string = "http://jdbc.datahunter.cn/sql"
	Fitting_proxy_url string = "http://jdbc.datahunter.cn/fitting"

	EE_CFG P = P{}

	Cron *cron.Cron
	Cmap cmap.ConcurrentMap
	Db   *scribble.Driver
)
