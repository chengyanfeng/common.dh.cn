package controllers

import (
	"fmt"
	"net/url"
	"strings"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"common.dh.cn/models"
	"common.dh.cn/def"
	"common.dh.cn/utils"
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

func (c *BaseController) EchoJson(p utils.P) {
	c.Data["json"] = p
	c.ServeJSON()
}

func (c *BaseController) EchoJsonMsg(msg interface{}) {
	c.Data["json"] = utils.P{"code": 200, "msg": msg}
	c.ServeJSON()
}

func (c *BaseController) EchoJsonOk(msg ...interface{}) {
	if msg == nil {
		msg = []interface{}{"ok"}
	}
	c.Data["json"] = utils.P{"code": 200, "msg": msg[0]}
	c.ServeJSON()
}

func (c *BaseController) EchoJsonErr(msg ...interface{}) {
	out := ""
	if msg != nil {
		for _, v := range msg {
			out = utils.JoinStr(out, v)
		}
	}
	c.Data["json"] = utils.P{"code": def.GENERAL_ERR, "msg": out}
	c.ServeJSON()
}

func (c *BaseController) FormToP(keys ...string) (p utils.P) {
	p = utils.P{}
	referer := c.Ctx.Request.Header["Dhref"]
	if !utils.IsEmpty(referer) && len(referer) > 0 {
		u, err := url.Parse(utils.Replace(referer[0], []string{"#"}, ""))
		if err == nil {
			vs := u.Query()
			for k, v := range vs {
				p[k] = utils.ToString(v)
			}
		}
	}
	r := c.Ctx.Request
	r.ParseForm()
	for k, v := range r.Form {
		if len(keys) > 0 {
			if utils.InArray(k, keys) {
				utils.SetKv(p, k, v)
			}
		} else {
			utils.SetKv(p, k, v)
		}
	}
	delete(p, "auth")
	utils.Debug("FormToP", p)
	return
}

func (c *BaseController) HeadToP() (p utils.P) {
	p = utils.P{}
	referer := c.Ctx.Request.Header["Dhref"]
	if !utils.IsEmpty(referer) && len(referer) > 0 {
		addr := referer[0]
		addr, _ = url.QueryUnescape(addr)
		if len(addr) > 1 {
			u, err := url.Parse("?" + addr)
			if err == nil {
				vs := u.Query()
				for k, v := range vs {
					p[k] = utils.ToString(v)
				}
			}
		}
		utils.Debug("HeadToP", p)
	}
	return
}

func (c *BaseController) PageParam() (start int, rows int) {
	page, _ := c.GetInt("page", 1)
	rows, _ = c.GetInt("rows", 10)
	start = (page - 1) * rows
	return
}

func (c *BaseController) GetOid(str string) bson.ObjectId {
	return utils.ToOid(c.GetString(str))
}

func (c *BaseController) Hostname() string {
	hostname := utils.ToString(c.Ctx.Request.Header.Get("Hostname"), "localhost:8080")
	return hostname
}

func (c *BaseController) LocalHost() string {
	return utils.JoinStr("http://localhost:", beego.BConfig.Listen.HTTPPort)
}

func (c *BaseController) Require(k ...string) {
	for _, v := range k {
		if utils.IsEmpty(c.GetString(v)) {
			c.EchoJsonErr(fmt.Sprintf("需要%v参数", v))
			c.StopRun()
		}
	}
}

func (c *BaseController) RequireOid(k ...string) {
	for _, v := range k {
		if !utils.IsOid(c.GetString(v)) {
			c.EchoJsonErr(fmt.Sprintf("%v参数必须是有效id", v))
			c.StopRun()
		}
	}
}

func (c *BaseController) HeadHref(str string) bool {
	referer := c.Ctx.Request.Header["Dhref"]
	if !utils.IsEmpty(referer) && len(referer) > 0 {
		addr := referer[0]
		addr, _ = url.QueryUnescape(addr)
		if strings.Contains(addr, str) {
			return true
		}
	}

	return false
}

func (c *BaseController) GetAuthUser() *models.DhUser {
	auth := c.GetString("auth")
	if auth == "" {
		auth = c.Ctx.GetCookie("auth")
	}
	if !utils.IsEmpty(auth) {
		return new(models.DhUser).Find("auth",auth)
	} else {
		return user
	}
}

func GetUserByEmail(email string) *P {
	user := D(User).Find(P{"email": email}).One()

	return user
}

func GetUserByAuth(auth string) *P {
	return D(User).Find(P{"auth": auth}).One()
}

func GetUserById(uid interface{}) *P {
	var oid bson.ObjectId
	switch uid.(type) {
	case string:
		oid = ToOid(uid.(string))
	case bson.ObjectId:
		oid = uid.(bson.ObjectId)
	}
}

func (c *BaseController) GetUserCorps(user_id string) []utils.P {
	corps := []utils.P{}
	//私人空间
	info := utils.P{}
	info["_id"] = user_id
	info["name"] = "私人空间"
	info["role"] = "admin"
	corps = append(corps,info)
	//其他团队
	filters := map[string]interface{}{}
	filters["user_id"] = user_id
	user_corps := new(models.DhUserCorp).List(filters)
	for _, v := range user_corps {
		corp := new(models.DhCorp).Find(v.CropId)
		info := utils.P{}
		info["_id"] = corp.ObjectId
		info["name"] = corp.Name
		info["role"] = v.Role
		corps = append(corps,info)
	}
	return corps
}
