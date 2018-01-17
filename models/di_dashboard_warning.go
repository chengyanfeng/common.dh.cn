package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiDashboardWarning struct {
	DiBase
	Id         int64     `json:"-"`
	ObjectId   string    `json:"_id"`
	WidgetId   string    `json:"widget_id"`
	UserId     string    `json:"user_id"`
	Sql        string    `json:"sql"`
	Field      string    `json:"field"`
	Key        string    `json:"key"`
	Compare    string    `json:"compare"`
	Value      int64     `json:"value"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiDashboardWarning))
}

func (m *DiDashboardWarning) TableName() string {
	return "di_dashboard_warning"
}

func (m *DiDashboardWarning) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiDashboardWarning) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiDashboardWarning) Find(args ...interface{}) *DiDashboardWarning {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiDashboardWarning)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiDashboardWarning) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiDashboardWarning) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiDashboardWarning) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiDashboardWarning) List(filters map[string]interface{}) []*DiDashboardWarning {
	var list []*DiDashboardWarning
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDashboardWarning) OrderList(filters map[string]interface{}, order ...string) []*DiDashboardWarning {
	var list []*DiDashboardWarning
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDashboardWarning) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiDashboardWarning) {
	var list []*DiDashboardWarning
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiDashboardWarning) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiDashboardWarning) {
	var list []*DiDashboardWarning
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
