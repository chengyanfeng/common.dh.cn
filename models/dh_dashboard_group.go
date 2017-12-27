package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhDashboardGroup struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    	Name string `json:"name"`
    	Status int `json:"status"`
    	Sort int `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhDashboardGroup))
}

func (m *DhDashboardGroup) TableName() string {
    return "dh_dashboard_group"
}

func (m *DhDashboardGroup) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhDashboardGroup) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhDashboardGroup) Find(args ...interface{}) *DhDashboardGroup {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhDashboardGroup)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhDashboardGroup) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DhDashboardGroup) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DhDashboardGroup) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *DhDashboardGroup) List(filters map[string]interface{}) []*DhDashboardGroup {
	var list []*DhDashboardGroup
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhDashboardGroup) OrderList(filters map[string]interface{},order ...string) []*DhDashboardGroup {
	var list []*DhDashboardGroup
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhDashboardGroup) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhDashboardGroup) {
	var list []*DhDashboardGroup
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}