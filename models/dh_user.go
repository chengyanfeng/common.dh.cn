package models

import (
	"time"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type DhUser struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    	Name string `json:"name"`
    	Email string `json:"email"`
    	Mobile string `json:"mobile"`
    	Password string `json:"password"`
    	Auth string `json:"auth"`
    	Avatar string `json:"avatar"`
    	Icode string `json:"icode"`
    	Status int `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhUser))
}

func (m *DhUser) TableName() string {
    return "dh_user"
}

func (m *DhUser) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhUser) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhUser) Find(args ...interface{}) *DhUser {
	data := m.find(m,args...)
	if data != nil {
		return data.(*DhUser)
	} else {
		return nil
	}
}

func (m *DhUser) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *DhUser) List(filters map[string]interface{}) []*DhUser {
	var list []*DhUser
	_, err := m.findByFilters(m,filters).All(&list)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return list
}

func (m *DhUser) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhUser) {
	var list []*DhUser
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m,page, page_size, filters).All(&list)
	if err != nil {
		utils.Error(err)
		return 0,0,nil
	}
	return total, total_page, list
}