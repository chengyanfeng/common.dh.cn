package def

import (
	"time"
)

const (
	DEFAULT_HTTP_TIMEOUT time.Duration = 60 * time.Second
	CACHE_TTL_DEFAULT                  = 60

	MODE_TEST string = "TEST"
	MODE_EE   string = "EE"

	AUDIT_WAIT string = "wait"
	AUDIT_PASS string = "pass"
	AUDIT_BAN  string = "ban"
)


var	SHARE_TYPE = map[string]int{
	"dh_dashboard_group":1,
	"dh_dashboard":1,
	"dh_storyboard_group":1,
	"dh_storyboard":1,
	"dh_datasource_group":1,
	"dh_datasource":1,
}
