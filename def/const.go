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
