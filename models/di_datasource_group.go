package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiDatasourceGroup struct {
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
	orm.RegisterModel(new(DiDatasourceGroup))
}

func (m *DiDatasourceGroup) TableName() string {
	return "di_datasource_group"
}

func (m *DiDatasourceGroup) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiDatasourceGroup) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiDatasourceGroup) Find(args ...interface{}) *DiDatasourceGroup {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiDatasourceGroup)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiDatasourceGroup) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiDatasourceGroup) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiDatasourceGroup) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiDatasourceGroup) List(filters map[string]interface{}) []*DiDatasourceGroup {
	var list []*DiDatasourceGroup
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDatasourceGroup) OrderList(filters map[string]interface{}, order ...string) []*DiDatasourceGroup {
	var list []*DiDatasourceGroup
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDatasourceGroup) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiDatasourceGroup) {
	var list []*DiDatasourceGroup
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiDatasourceGroup) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiDatasourceGroup) {
	var list []*DiDatasourceGroup
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
