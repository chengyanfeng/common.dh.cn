package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiDashboardGroup struct {
	DiBase
	Id         int64     `json:"-"`
	ObjectId   string    `json:"_id"`
	Name       string    `json:"name"`
	Status     int       `json:"status"`
	Sort       int       `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiDashboardGroup))
}

func (m *DiDashboardGroup) TableName() string {
	return "di_dashboard_group"
}

func (m *DiDashboardGroup) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiDashboardGroup) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiDashboardGroup) Find(args ...interface{}) *DiDashboardGroup {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiDashboardGroup)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiDashboardGroup) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiDashboardGroup) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiDashboardGroup) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiDashboardGroup) List(filters map[string]interface{}) []*DiDashboardGroup {
	var list []*DiDashboardGroup
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDashboardGroup) OrderList(filters map[string]interface{}, order ...string) []*DiDashboardGroup {
	var list []*DiDashboardGroup
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDashboardGroup) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiDashboardGroup) {
	var list []*DiDashboardGroup
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiDashboardGroup) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiDashboardGroup) {
	var list []*DiDashboardGroup
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
