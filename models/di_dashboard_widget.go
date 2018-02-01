package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiDashboardWidget struct {
	DiBase
	Id           int64     `json:"-"`
	ObjectId     string    `json:"_id"`
	DashboardId  string    `json:"dashboard_id"`
	DatasourceId string    `json:"datasource_id"`
	Grid         string    `json:"grid"`
	Config       string    `json:"config"`
	Filter       int       `json:"filter"`
	Status       int       `json:"status"`
	CreateTime   time.Time `json:"-"`
	UpdateTime   time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiDashboardWidget))
}

func (m *DiDashboardWidget) TableName() string {
	return "di_dashboard_widget"
}

func (m *DiDashboardWidget) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiDashboardWidget) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiDashboardWidget) Find(args ...interface{}) *DiDashboardWidget {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiDashboardWidget)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiDashboardWidget) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiDashboardWidget) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiDashboardWidget) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiDashboardWidget) List(filters map[string]interface{}) []*DiDashboardWidget {
	var list []*DiDashboardWidget
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDashboardWidget) OrderList(filters map[string]interface{}, order ...string) []*DiDashboardWidget {
	var list []*DiDashboardWidget
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDashboardWidget) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiDashboardWidget) {
	var list []*DiDashboardWidget
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiDashboardWidget) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiDashboardWidget) {
	var list []*DiDashboardWidget
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
