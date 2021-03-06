package utils

import (
	"encoding/json"

	"github.com/clbanning/mxj"
)

func JsonDecode(b []byte) (p *P) {
	p = &P{}
	err := json.Unmarshal(b, p)
	if err != nil {
		Error("JsonDecode", string(b), err)
	}
	return
}

func JsonEncode(v interface{}) (r string) {
	b, err := json.Marshal(v)
	if err != nil {
		Error(err)
	}
	r = string(b)
	return
}

func IsJson(b []byte) bool {
	var j json.RawMessage
	return json.Unmarshal(b, &j) == nil
}

func JsonDecodeArray(b []byte) (p []P, e error) {
	p = []P{}
	e = json.Unmarshal(b, &p)
	return
}

func JsonDecodeArray_str(b []byte) (p []string, e error) {
	p = []string{}
	e = json.Unmarshal(b, &p)
	return
}

func JsonDecodeArrays(b []byte) (p *[]P) {
	p = &[]P{}
	e := json.Unmarshal(b, p)
	if e != nil {
		Error(e)
	}
	return
}

func JsonDecodeStrings(s string) (r []string) {
	r = []string{}
	e := json.Unmarshal([]byte(s), &r)
	if e != nil {
		Error(e, s)
	}
	return
}

func JoinStr(val ...interface{}) (r string) {
	for _, v := range val {
		r += ToString(v)
	}
	return
}

func Xml2Json(src string) (s string, err error) {
	m, err := mxj.NewMapXml([]byte(src))
	return JsonEncode(m), err
}
