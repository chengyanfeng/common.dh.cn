package models

import (
	"time"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type DhCorp struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Mobile string  `json:"mobile"`
	Vcode string `json:"vcode"`
	Status int `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhCorp))
}

func (m *DhCorp) TableName() string {
    return "dh_user"
}

func (m *DhCorp) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhCorp) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhCorp) Find(args ...interface{}) *DhCorp {
	data := m.find(m,args...)
	if data != nil {
		return data.(*DhCorp)
	} else {
		return nil
	}
}

func (m *DhCorp) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *DhCorp) List(filters map[string]interface{}) []*DhCorp {
	var list []*DhCorp
	_, err := m.filter(m,filters).All(&list)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return list
}

func (m *DhCorp) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, result []*DhCorp) {
	var list []*DhCorp
	total = m.count(m, filters)
	_, err := m.pager(m,page, page_size, filters).All(&list)
	if err != nil {
		utils.Error(err)
		return 0,nil
	}
	return total,list
}
