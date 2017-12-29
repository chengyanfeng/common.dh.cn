package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhDatasourceGroup struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    	Name string `json:"name"`
    	Status int `json:"status"`
    	Sort int `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhDatasourceGroup))
}

func (m *DhDatasourceGroup) TableName() string {
    return "dh_datasource_group"
}

func (m *DhDatasourceGroup) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhDatasourceGroup) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhDatasourceGroup) Find(args ...interface{}) *DhDatasourceGroup {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhDatasourceGroup)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhDatasourceGroup) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DhDatasourceGroup) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DhDatasourceGroup) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *DhDatasourceGroup) List(filters map[string]interface{}) []*DhDatasourceGroup {
	var list []*DhDatasourceGroup
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhDatasourceGroup) OrderList(filters map[string]interface{}, order ...string) []*DhDatasourceGroup {
	var list []*DhDatasourceGroup
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhDatasourceGroup) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhDatasourceGroup) {
	var list []*DhDatasourceGroup
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}

func (m *DhDatasourceGroup) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DhDatasourceGroup) {
	var list []*DhDatasourceGroup
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}