package controllers

import (
	"fmt"
	"reflect"
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

func (c *BaseController) GetAuthUser() (m *models.DhUser) {
	defer func() {
		if m == nil{
			c.EchoJsonErr("用户不存在")
			c.StopRun()
		}
	}()
	auth := c.GetString("auth")
	if auth == "" {
		auth = c.Ctx.GetCookie("auth")
	}
	if !utils.IsEmpty(auth) {
		user := new(models.DhUser).Find("auth", auth)
		if (user != nil) {
			if user.Status != -1 {
				m = user
			} else {
				m = nil
			}
		} else {
			m = nil
		}
	} else {
		m = nil
	}
	return
}

func (c *BaseController) GetUserCorps(user_id string) []utils.P {
	corps := []utils.P{}
	//私人空间
	info := utils.P{}
	info["_id"] = user_id
	info["name"] = "私人空间"
	info["role"] = "admin"
	corps = append(corps, info)
	//其他团队
	filters := map[string]interface{}{}
	filters["user_id"] = user_id
	user_corps := new(models.DhUserCorp).List(filters)
	for _, v := range user_corps {
		corp := new(models.DhCorp).Find(v.CorpId)
		info := utils.P{}
		info["_id"] = corp.ObjectId
		info["name"] = corp.Name
		info["role"] = v.Role
		corps = append(corps, info)
	}
	return corps
}

func (c *BaseController) Notify(from_crop_id string, from_user_id string, user_id string, notify_type string, config interface{}) {
	notify := new(models.DhNotify)
	notify.FromCropId = from_crop_id
	notify.FromUserId = from_user_id
	notify.UserId = user_id
	notify.Type = notify_type
	notify.Config = utils.JsonEncode(config)
	result := notify.Save()
	if result {
		//TODO Websocket NotifySend
	}
}

func (c *BaseController) SaveRelation(id int64, object_id string, crop_id string, user_id string, relate_type string, relate_id string, name string, auth string) bool {
	relation := new(models.DhRelation)
	if id != 0 {
		relation.Id = id
		relation.ObjectId = object_id
	}
	relation.CorpId = crop_id
	relation.UserId = user_id
	relation.RelateType = relate_type
	relation.RelateId = relate_id
	relation.Name = name
	relation.Auth = auth
	relation.Sort = 0
	return relation.Save()
}

func (c *BaseController) SortRelation(crop_id string, user_id string, relate_type string, relate_ids []string) bool {
	o := new(models.DhBase).Orm()
	for k, relate_id := range relate_ids {
		params := map[string]interface{}{}
		params["crop_id"] = crop_id
		params["user_id"] = user_id
		params["relate_type"] = relate_type
		params["relate_id"] = relate_id
		relation := new(models.DhRelation).Find(params)
		if relation != nil {
			relation.Sort = k
			result := relation.Save()
			if !result {
				o.Rollback()
				return false
			}
		}
	}
	o.Commit()
	return true
}

func (c *BaseController) Share(crop_id string, user_ids []string, relate_type string, relate_id string) bool {
	o := new(models.DhBase).Orm()
	var err error
	err = o.Begin()
	if err != nil {
		return false
	}
	share_name := c.GetShareName(relate_type, relate_id)
	if share_name == "" {
		fmt.Println(2)
		o.Commit()
		return false
	}
	result := c.RemoveShare(relate_type, relate_id, "share")
	if !result {
		fmt.Println(3)
		o.Rollback()
		return false
	}
	for _, user_id := range user_ids {
		result := c.SaveRelation(0, "", crop_id, user_id, relate_type, relate_id, share_name, "share")
		if !result {
			fmt.Println(4)
			o.Rollback()
			return false
		}
	}
	o.Commit()
	return true
}

func (c *BaseController) ShareOut(user_emails []string, relate_type string, relate_id string) bool {
	user_ids := []string{}
	for _, user_email := range user_emails {
		user := new(models.DhUser).Find("email", user_email)
		if user != nil {
			user_ids = append(user_ids, user.ObjectId)
		}
	}
	if len(user_ids) > 0 {
		o := new(models.DhBase).Orm()
		var err error
		err = o.Begin()
		if err != nil {
			return false
		}
		share_name := c.GetShareName(relate_type, relate_id)
		if share_name == "" {
			o.Commit()
			return false
		}
		result := c.RemoveShare(relate_type, relate_id, "share_out")
		if !result {
			o.Rollback()
			return false
		}
		for _, user_id := range user_ids {
			//跨组分享进入默认分组
			result := c.SaveRelation(0, "", user_id, user_id, relate_type, relate_id, share_name, "share_out")
			if !result {
				o.Rollback()
				return false
			}
		}
		o.Commit()
		return true
	} else {
		return false
	}
}

func (c *BaseController) RemoveShare(relate_type string, relate_id string, auth string) bool {
	params := map[string]interface{}{}
	params["relate_type"] = relate_type
	params["relate_id"] = relate_id
	params["auth"] = auth
	return new(models.DhRelation).Delete(params)
}

func (c *BaseController) GetShareName(relate_type string, relate_id string) string {
	var relate_object interface{}
	switch (relate_type) {
	case "dh_dashboard_group":
		relate_object = new(models.DhDashboardGroup).Find(relate_id)
	case "dh_dashboard":
		relate_object = new(models.DhDashboard).Find(relate_id)
	case "dh_storyboard_group":
		relate_object = new(models.DhStoryboardGroup).Find(relate_id)
	case "dh_storyboard":
		relate_object = new(models.DhStoryboard).Find(relate_id)
	case "dh_datasource_group":
		relate_object = new(models.DhDatasourceGroup).Find(relate_id)
	case "dh_datasource":
		relate_object = new(models.DhDatasource).Find(relate_id)
	}
	if relate_object == nil {
		return ""
	} else {
		return reflect.ValueOf(relate_object).Elem().FieldByName("Name").String()
	}
}

func (c *BaseController) GetShareUsers(relate_type string, relate_id string) []*models.DhUser {
	users := []*models.DhUser{}
	params := map[string]interface{}{}
	params["relate_type"] = relate_type
	params["relate_id"] = relate_id
	params["auth"] = "share"
	relations := new(models.DhRelation).List(params)
	for _, v := range relations {
		user := new(models.DhUser).Find(v.UserId)
		if user != nil {
			users = append(users, user)
		}
	}
	return users
}

func (c *BaseController) GetShareOutUsers(relate_type string, relate_id string) []*models.DhUser {
	users := []*models.DhUser{}
	params := map[string]interface{}{}
	params["relate_type"] = relate_type
	params["relate_id"] = relate_id
	params["auth"] = "share_out"
	relations := new(models.DhRelation).List(params)
	for _, v := range relations {
		user := new(models.DhUser).Find(v.UserId)
		if user != nil {
			users = append(users, user)
		}
	}
	return users
}
