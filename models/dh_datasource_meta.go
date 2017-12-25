package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhDatasourceMeta struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    	DatasourceId string `json:"datasource_id"`
    	Column string `json:"column"`
    	Name string `json:"name"`
    	Type string `json:"type"`
    	Extra string `json:"extra"`
    	Status int `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhDatasourceMeta))
}

func (m *DhDatasourceMeta) TableName() string {
    return "dh_datasource_meta"
}

func (m *DhDatasourceMeta) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhDatasourceMeta) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhDatasourceMeta) Find(args ...interface{}) *DhDatasourceMeta {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhDatasourceMeta)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhDatasourceMeta) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *DhDatasourceMeta) SoftDelete(index interface{}) bool {
	return m.softDelete(m,index)
}

func (m *DhDatasourceMeta) List(filters map[string]interface{}) []*DhDatasourceMeta {
	var list []*DhDatasourceMeta
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhDatasourceMeta) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhDatasourceMeta) {
	var list []*DhDatasourceMeta
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}