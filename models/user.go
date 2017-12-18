package models

import (
	"time"
	"common.dh.cn/util"
	"github.com/rs/xid"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id string `orm:"pk"`
	Email string 
	Mobile string 
	Password string 
	Auth string
	Name string
	CreateTime time.Time
	UpdateTime time.Time
}


func (u *User) TableName() string {
    return "dh_user"
}

func (this *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *User) FindOne(id string) *User {
	err := this.Query().Filter("id", id).One(this)
	if err != nil {
		return nil
	}
	return this
}

func (this *User) Create() bool {
	this.Id = xid.New().String()
	this.CreateTime = time.Now()
	this.UpdateTime = time.Now()
	_, err := orm.NewOrm().Insert(this)
	if err != nil && err.Error() != "no LastInsertId available" {
		util.Error(err)
		return false
	} else {
		return true
	}
}

func (this *User) Update() bool {
	this.UpdateTime = time.Now()
	_, err := orm.NewOrm().Update(this)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (this *User) Delete(id string) bool {
	num, err := this.Query().Filter("id", id).Delete()
	if err != nil {
		return false
	}
	return num >= 1
}