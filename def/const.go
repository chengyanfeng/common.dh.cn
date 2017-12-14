package def

import (
	"time"
)

const (
	DEFAULT_HTTP_TIMEOUT time.Duration = 60 * time.Second
	CACHE_TTL_DEFAULT                  = 60

	GENERAL_ERR   int = 400
	NO_SHARE_ERR  int = 401
	PWD_REQUIRED  int = 402
	IP_LIMIT_ERR  int = 403
	NOT_LOGIN     int = 404
	DATA_LINE_ERR int = 405

	MODE_TEST string = "TEST"
	MODE_EE   string = "EE"

	Cname string = "job"

	// Collection
	BoardGroup string = "boardgroup"
	Corp       string = "corp"
	Dashboard  string = "dashboard"
	DataSource string = "datasource"
	Record     string = "record"
	DbConn     string = "dbconn"
	DbPos      string = "dbpos"
	DsGroup    string = "dsgroup"
	DsType     string = "dstype"
	Icode      string = "icode"
	Relation   string = "relat"
	Relation14 string = "relat14"
	Snapshot   string = "snapshot"
	User       string = "user"
	UserCorp   string = "usercorp"
	WxUser     string = "wxuser"
	WhiteList  string = "whitelist"
	Widget     string = "widget"
	Admin      string = "admin"
	Root       string = "root"
	ChatMsg    string = "chatmsg"
	Linkshare  string = "linkshare"
	ApiData    string = "apidata"
	Comment    string = "comment"
	Notify     string = "notify"
	Warning    string = "warning"
)
