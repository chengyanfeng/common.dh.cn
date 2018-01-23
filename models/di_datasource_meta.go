package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiDatasourceMeta struct {
	DiBase
	Id           int64     `json:"-"`
	ObjectId     string    `json:"_id"`
	DatasourceId string    `json:"datasource_id"`
	Column       string    `json:"column"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	IsCalc       int       `json:"is_calc"`
	Extra        string    `json:"extra"`
	Status       int       `json:"status"`
	CreateTime   time.Time `json:"-"`
	UpdateTime   time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiDatasourceMeta))
}

func (m *DiDatasourceMeta) TableName() string {
	return "di_datasource_meta"
}

func (m *DiDatasourceMeta) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiDatasourceMeta) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiDatasourceMeta) Find(args ...interface{}) *DiDatasourceMeta {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiDatasourceMeta)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiDatasourceMeta) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiDatasourceMeta) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiDatasourceMeta) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiDatasourceMeta) List(filters map[string]interface{}) []*DiDatasourceMeta {
	var list []*DiDatasourceMeta
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDatasourceMeta) OrderList(filters map[string]interface{}, order ...string) []*DiDatasourceMeta {
	var list []*DiDatasourceMeta
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDatasourceMeta) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiDatasourceMeta) {
	var list []*DiDatasourceMeta
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiDatasourceMeta) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiDatasourceMeta) {
	var list []*DiDatasourceMeta
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
