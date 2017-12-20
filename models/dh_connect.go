package models

import (
	"time"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type DhConnect struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    	CropId string `json:"crop_id"`
    	UserId string `json:"user_id"`
    	Name string `json:"name"`
    	Type string `json:"type"`
    	Config string `json:"config"`
    	Status int `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhConnect))
}

func (m *DhConnect) TableName() string {
    return "dh_connect"
}

func (m *DhConnect) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhConnect) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhConnect) Find(args ...interface{}) *DhConnect {
	data := m.find(m,args...)
	if data != nil {
		return data.(*DhConnect)
	} else {
		return nil
	}
}

func (m *DhConnect) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *DhConnect) List(filters map[string]interface{}) []*DhConnect {
	var list []*DhConnect
	_, err := m.findByFilters(m,filters).All(&list)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return list
}

func (m *DhConnect) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhConnect) {
	var list []*DhConnect
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m,page, page_size, filters).All(&list)
	if err != nil {
		utils.Error(err)
		return 0,0,nil
	}
	return total, total_page, list
}