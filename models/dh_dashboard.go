package models

import (
	"time"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type DhDashboard struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    GroupId string `json:"group_id"`
    Name string `json:"name"`
    Thumbnail string `json:"thumbnail"`
    Status int `json:"status"`
    Sort int `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhDashboard))
}

func (m *DhDashboard) TableName() string {
    return "dh_dashboard"
}

func (m *DhDashboard) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhDashboard) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhDashboard) Find(args ...interface{}) *DhDashboard {
	data := m.find(m,args...)
	if data != nil {
		return data.(*DhDashboard)
	} else {
		return nil
	}
}

func (m *DhDashboard) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *DhDashboard) List(filters map[string]interface{}) []*DhDashboard {
	var list []*DhDashboard
	_, err := m.findByFilters(m,filters).All(&list)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return list
}

func (m *DhDashboard) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhDashboard) {
	var list []*DhDashboard
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m,page, page_size, filters).All(&list)
	if err != nil {
		utils.Error(err)
		return 0,0,nil
	}
	return total, total_page, list
}