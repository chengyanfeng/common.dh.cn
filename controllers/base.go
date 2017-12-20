package controllers

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"

	. "common.dh.cn/def"
	. "common.dh.cn/models"
	. "common.dh.cn/utils"
)

var Num = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Echo(msg ...interface{}) {
	var out string = ""
	for _, v := range msg {
		out += fmt.Sprintf("%v", v)
	}

	c.Ctx.WriteString(out)
}

func (c *BaseController) EchoJson(p P) {
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.WriteString(JsonEncode(p))
}

func (c *BaseController) EchoJsonMsg(msg interface{}) {
	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.WriteString(JsonEncode(P{"code": 200, "msg": msg}))
}

func (c *BaseController) EchoJsonOk(msg ...interface{}) {
	if msg == nil {
		msg = []interface{}{"ok"}
	}

	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.WriteString(JsonEncode(P{"code": 200, "msg": msg[0]}))
}

func (c *BaseController) EchoJsonErr(msg ...interface{}) {
	out := ""
	if msg != nil {
		for _, v := range msg {
			out = JoinStr(out, v)
		}
	}

	c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	c.Ctx.WriteString(JsonEncode(P{"code": GENERAL_ERR, "msg": out}))
}

// 把form表单内容赋予P结构体
func (c *BaseController) FormToP(keys ...string) (p P) {
	p = P{}
	referer := c.Ctx.Request.Header["Dhref"]
	if !IsEmpty(referer) && len(referer) > 0 {
		u, err := url.Parse(Replace(referer[0], []string{"#"}, ""))
		if err == nil {
			vs := u.Query()
			for k, v := range vs {
				p[k] = ToString(v)
			}
		}
	}

	r := c.Ctx.Request
	r.ParseForm()
	for k, v := range r.Form {
		if len(keys) > 0 {
			if InArray(k, keys) {
				setKv(p, k, v)
			}
		} else {
			setKv(p, k, v)
		}
	}

	delete(p, "auth")

	return
}

func setKv(p P, k string, v []string) {
	if len(v) == 1 {
		if len(v[0]) > 0 {
			p[k] = v[0]
		}
	} else {
		p[k] = v
	}
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func (c *BaseController) PageParam() (start int, rows int) {
	page, _ := c.GetInt("page", 1)
	rows, _ = c.GetInt("rows", 10)
	start = (page - 1) * rows

	return
}

func (c *BaseController) GetOid(str string) (oid bson.ObjectId) {
	oid = ToOid(c.GetString(str))

	return
}

func (c *BaseController) Hostname() string {
	hostname := ToString(c.Ctx.Request.Header.Get("Hostname"), "localhost:8080")

	return hostname
	//return c.Ctx.Request.Host
}

func (c *BaseController) LocalHost() string {
	return JoinStr("http://localhost:", beego.BConfig.Listen.HTTPPort)
}

func (c *BaseController) Require(k ...string) {
	for _, v := range k {
		if IsEmpty(c.GetString(v)) {
			c.EchoJsonErr(fmt.Sprintf("需要%v参数", v))
			c.StopRun()
		}
	}
}

func (c *BaseController) RequireOid(k ...string) {
	for _, v := range k {
		if !IsOid(c.GetString(v)) {
			c.EchoJsonErr(fmt.Sprintf("%v参数必须是有效id", v))
			c.StopRun()
		}
	}
}

func (c *BaseController) HeadToP() (p P) {
	p = P{}
	referer := c.Ctx.Request.Header["Dhref"]
	if !IsEmpty(referer) && len(referer) > 0 {
		addr := referer[0]
		addr, _ = url.QueryUnescape(addr)
		if len(addr) > 1 {
			u, err := url.Parse("?" + addr)
			if err == nil {
				vs := u.Query()
				for k, v := range vs {
					p[k] = ToString(v)
				}
			}
		}
		Debug("HeadToP", p)
	}

	return
}

func (c *BaseController) HeadHref(str string) bool {
	referer := c.Ctx.Request.Header["Dhref"]
	if !IsEmpty(referer) && len(referer) > 0 {
		addr := referer[0]
		addr, _ = url.QueryUnescape(addr)
		if strings.Contains(addr, str) {
			return true
		}
	}

	return false
}

func (c *BaseController) GetAuthUser() P {
	auth := c.GetString("auth")
	if auth == "" {
		auth = c.Ctx.GetCookie("auth")
	}

	user := P{}
	if !IsEmpty(auth) {
		user = *GetUserByAuth(auth)
	}

	if user["_id"] == nil {
		return nil
	} else {
		return user
	}
}

func GetUserByEmail(email string) *P {
	user := D(User).Find(P{"email": email}).One()

	return user
}

func GetUserByAuth(auth string) *P {
	return D(User).Find(P{"auth": auth}).Cache().One()
}

func GetUserById(uid interface{}) *P {
	var oid bson.ObjectId
	switch uid.(type) {
	case string:
		oid = ToOid(uid.(string))
	case bson.ObjectId:
		oid = uid.(bson.ObjectId)
	}
	user := D(User).Find(P{"_id": oid}).Cache().One()
	return user
}
