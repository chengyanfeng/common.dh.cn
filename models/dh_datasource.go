package models

import (
	"time"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type DhDatasource struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    GroupId string `json:"group_id"`
    Name string `json:"name"`
    Table string `json:"table"`
    Type string `json:"type"`
    Format string `json:"format"`
    Url string `json:"url"`
    ConnectId string `json:"connect_id"`
    Status int `json:"status"`
    Sort int `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhDatasource))
}

func (m *DhDatasource) TableName() string {
    return "dh_datasource"
}

func (m *DhDatasource) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhDatasource) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhDatasource) Find(args ...interface{}) *DhDatasource {
	data := m.find(m,args...)
	if data != nil {
		return data.(*DhDatasource)
	} else {
		return nil
	}
}

func (m *DhDatasource) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *DhDatasource) List(filters map[string]interface{}) []*DhDatasource {
	var list []*DhDatasource
	_, err := m.findByFilters(m,filters).All(&list)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return list
}

func (m *DhDatasource) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhDatasource) {
	var list []*DhDatasource
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m,page, page_size, filters).All(&list)
	if err != nil {
		utils.Error(err)
		return 0,0,nil
	}
	return total, total_page, list
}