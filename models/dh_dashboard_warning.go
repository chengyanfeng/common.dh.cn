package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhDashboardWarning struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    	WidgetId string `json:"widget_id"`
    	Key string `json:"key"`
    	Compare string `json:"compare"`
    	Value int64 `json:"value"`
    	Status int `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhDashboardWarning))
}

func (m *DhDashboardWarning) TableName() string {
    return "dh_dashboard_warning"
}

func (m *DhDashboardWarning) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhDashboardWarning) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhDashboardWarning) Find(args ...interface{}) *DhDashboardWarning {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhDashboardWarning)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhDashboardWarning) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *DhDashboardWarning) SoftDelete(index interface{}) bool {
	return m.softDelete(m,index)
}

func (m *DhDashboardWarning) List(filters map[string]interface{}) []*DhDashboardWarning {
	var list []*DhDashboardWarning
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhDashboardWarning) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhDashboardWarning) {
	var list []*DhDashboardWarning
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}