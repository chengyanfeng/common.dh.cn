package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiDatasourceRelation struct {
	DiBase
	Id          int64     `json:"-"`
	ObjectId    string    `json:"_id"`
	CorpId      string    `json:"corp_id"`
	UserId      string    `json:"user_id"`
	Datasource1 string    `json:"datasource1"`
	Datasource2 string    `json:"datasource2"`
	Join        string    `json:"join"`
	On          string    `json:"on"`
	Status      int       `json:"status"`
	CreateTime  time.Time `json:"-"`
	UpdateTime  time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiDatasourceRelation))
}

func (m *DiDatasourceRelation) TableName() string {
	return "di_datasource_relation"
}

func (m *DiDatasourceRelation) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiDatasourceRelation) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiDatasourceRelation) Find(args ...interface{}) *DiDatasourceRelation {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiDatasourceRelation)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiDatasourceRelation) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiDatasourceRelation) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiDatasourceRelation) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiDatasourceRelation) List(filters map[string]interface{}) []*DiDatasourceRelation {
	var list []*DiDatasourceRelation
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDatasourceRelation) OrderList(filters map[string]interface{}, order ...string) []*DiDatasourceRelation {
	var list []*DiDatasourceRelation
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDatasourceRelation) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiDatasourceRelation) {
	var list []*DiDatasourceRelation
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiDatasourceRelation) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiDatasourceRelation) {
	var list []*DiDatasourceRelation
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
