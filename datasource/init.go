package datasource

import (
	"os"
	"strings"

	. "common.dh.cn/def"
	. "common.dh.cn/util"
)

func Init() {
	MODE = strings.ToUpper(Trim(os.Getenv("mode")))
	MONGO_HOST = Trim(os.Getenv("mongo"))
	if MODE == MODE_EE {
		EE = true
		Jdbc_proxy_url = "http://127.0.0.1:4567/sql"
		EE_MONGO["host"] = "127.0.0.1"
	}
	if IsEmpty(MONGO_HOST) {
		MONGO_HOST = "mongo_db"
	}
	DH_MONGO["host"] = MONGO_HOST
	if MODE == MODE_TEST {
		DH_MONGO["name"] = "dhtest"
		Jdbc_proxy_url = "http://jdbcdev.datahunter.cn/sql"
	}
}
