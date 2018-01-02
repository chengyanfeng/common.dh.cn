package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhDatasourceRelation struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
	Datasource1 string `json:"datasource1"`
	Datasource2 string `json:"datasource2"`
	Join string `json:"join"`
	On string `json:"on"`
	Status int `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhDatasourceRelation))
}

func (m *DhDatasourceRelation) TableName() string {
    return "dh_datasource_relation"
}

func (m *DhDatasourceRelation) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhDatasourceRelation) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhDatasourceRelation) Find(args ...interface{}) *DhDatasourceRelation {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhDatasourceRelation)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhDatasourceRelation) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DhDatasourceRelation) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DhDatasourceRelation) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *DhDatasourceRelation) List(filters map[string]interface{}) []*DhDatasourceRelation {
	var list []*DhDatasourceRelation
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhDatasourceRelation) OrderList(filters map[string]interface{}, order ...string) []*DhDatasourceRelation {
	var list []*DhDatasourceRelation
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhDatasourceRelation) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhDatasourceRelation) {
	var list []*DhDatasourceRelation
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}

func (m *DhDatasourceRelation) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DhDatasourceRelation) {
	var list []*DhDatasourceRelation
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}