package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhDatasourceModify struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
	UserId string `json:"user_id"`
	DatasourceId string `json:"datasource_id"`
	Table string `json:"table"`
	RowId string `json:"row_id"`
	Column string `json:"column"`
	Type string `json:"type"`
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhDatasourceModify))
}

func (m *DhDatasourceModify) TableName() string {
    return "dh_datasource_modify"
}

func (m *DhDatasourceModify) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhDatasourceModify) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhDatasourceModify) Find(args ...interface{}) *DhDatasourceModify {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhDatasourceModify)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhDatasourceModify) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DhDatasourceModify) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DhDatasourceModify) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *DhDatasourceModify) List(filters map[string]interface{}) []*DhDatasourceModify {
	var list []*DhDatasourceModify
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhDatasourceModify) OrderList(filters map[string]interface{}, order ...string) []*DhDatasourceModify {
	var list []*DhDatasourceModify
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhDatasourceModify) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhDatasourceModify) {
	var list []*DhDatasourceModify
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}

func (m *DhDatasourceModify) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DhDatasourceModify) {
	var list []*DhDatasourceModify
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}