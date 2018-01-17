package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiDashboard struct {
	DiBase
	Id         int64     `json:"-"`
	ObjectId   string    `json:"_id"`
	GroupId    string    `json:"group_id"`
	Name       string    `json:"name"`
	Thumbnail  string    `json:"thumbnail"`
	Status     int       `json:"status"`
	Sort       int       `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiDashboard))
}

func (m *DiDashboard) TableName() string {
	return "di_dashboard"
}

func (m *DiDashboard) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiDashboard) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiDashboard) Find(args ...interface{}) *DiDashboard {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiDashboard)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiDashboard) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiDashboard) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiDashboard) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiDashboard) List(filters map[string]interface{}) []*DiDashboard {
	var list []*DiDashboard
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDashboard) OrderList(filters map[string]interface{}, order ...string) []*DiDashboard {
	var list []*DiDashboard
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDashboard) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiDashboard) {
	var list []*DiDashboard
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiDashboard) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiDashboard) {
	var list []*DiDashboard
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
