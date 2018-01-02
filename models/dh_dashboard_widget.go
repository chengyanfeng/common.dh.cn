package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhDashboardWidget struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
	DashboardId string `json:"dashboard_id"`
	Grid string `json:"grid"`
	Config string `json:"config"`
	Status int `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhDashboardWidget))
}

func (m *DhDashboardWidget) TableName() string {
    return "dh_dashboard_widget"
}

func (m *DhDashboardWidget) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhDashboardWidget) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhDashboardWidget) Find(args ...interface{}) *DhDashboardWidget {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhDashboardWidget)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhDashboardWidget) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DhDashboardWidget) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DhDashboardWidget) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *DhDashboardWidget) List(filters map[string]interface{}) []*DhDashboardWidget {
	var list []*DhDashboardWidget
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhDashboardWidget) OrderList(filters map[string]interface{}, order ...string) []*DhDashboardWidget {
	var list []*DhDashboardWidget
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhDashboardWidget) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhDashboardWidget) {
	var list []*DhDashboardWidget
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}

func (m *DhDashboardWidget) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DhDashboardWidget) {
	var list []*DhDashboardWidget
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}