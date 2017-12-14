package def

import (
	"time"
)

const (
	DEFAULT_HTTP_TIMEOUT time.Duration = 60 * time.Second
	CACHE_TTL_DEFAULT                  = 60

	GENERAL_ERR int = 400

	MODE_TEST string = "TEST"
	MODE_EE   string = "EE"

	Cname string = "job"
)
