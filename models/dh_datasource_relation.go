package models

import (
	"time"
	"common.dh.cn/utils"
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
		return data.(*DhDatasourceRelation)
	} else {
		return nil
	}
}

func (m *DhDatasourceRelation) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *DhDatasourceRelation) List(filters map[string]interface{}) []*DhDatasourceRelation {
	var list []*DhDatasourceRelation
	_, err := m.findByFilters(m,filters).All(&list)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return list
}

func (m *DhDatasourceRelation) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhDatasourceRelation) {
	var list []*DhDatasourceRelation
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m,page, page_size, filters).All(&list)
	if err != nil {
		utils.Error(err)
		return 0,0,nil
	}
	return total, total_page, list
}