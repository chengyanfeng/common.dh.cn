package models

import (
	"time"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type DhApiData struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    UserId string `json:"user_id"`
    Type string `json:"type"`
    Content string `json:"content"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhApiData))
}

func (m *DhApiData) TableName() string {
    return "dh_api_data"
}

func (m *DhApiData) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhApiData) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhApiData) Find(args ...interface{}) *DhApiData {
	data := m.find(m,args...)
	if data != nil {
		return data.(*DhApiData)
	} else {
		return nil
	}
}

func (m *DhApiData) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *DhApiData) List(filters map[string]interface{}) []*DhApiData {
	var list []*DhApiData
	_, err := m.findByFilters(m,filters).All(&list)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return list
}

func (m *DhApiData) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhApiData) {
	var list []*DhApiData
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m,page, page_size, filters).All(&list)
	if err != nil {
		utils.Error(err)
		return 0,0,nil
	}
	return total, total_page, list
}