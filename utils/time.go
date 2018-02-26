package utils

import (
	"time"
)

func Timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func DateTimeStr() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

func ToDate(s string) (str string, e error) {
	fmt := []string{
		"2006-1-2 15:04:05",
		"2006-01-02T15:04:05",
		"2006/1/2 15:04:05",

		"2006/1/2",
		"2006-1-2",
		"2006.1.2",
		"1-2-2006",
		"1-2-06",
		"200601",
		"2006年1月",
		"2006年1月2日 15:04:05",
		"2006年1月2日"}
	var t time.Time
	for _, f := range fmt {
		t, e = time.Parse(f, s)
		if e == nil {
			return t.Format("2006-01-02 15:04:05"), e
		}
	}
	s = ""
	return s, e
}

func IsDate(s interface{}) bool {
	_, e := ToDate(ToString(s))
	return e == nil
}
