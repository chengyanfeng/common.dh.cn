package utils

import (
	"fmt"
	"time"

	"common.dh.cn/def"
	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
)

var JDBCLogger *logrus.Logger
var JDBCUrl string

func init() {
	JDBCLogger = GetLogger("jdbc")
	JDBCUrl = beego.AppConfig.DefaultString("jdbc_url", "http://jdbcdev.datahunter.cn/sql")
}

func JDBC(sql string, db P) (result string, err error) {
	db_config := JsonEncode(db)
	logger := JDBCLogger.WithFields(logrus.Fields{
		"sql": sql,
		"db":  db_config,
	})
	begin := time.Now()
	logger.Info("begin")
	result, err = HttpPost(def.Jdbc_proxy_url, nil, &P{"sql": sql, "db": db_config})
	if err != nil {
		logger.Error(err)
	}
	finish := time.Now()
	nanoseconds := finish.Sub(begin).Nanoseconds()
	milliseconds := fmt.Sprintf("%d.%d", nanoseconds/1e6, nanoseconds%1e6)
	logger.WithField("consume", milliseconds).Info("finish")
	return
}
