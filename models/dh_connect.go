package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhConnect struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    	CorpId string `json:"corp_id"`
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
		_data,ok := data.(*DhConnect)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhConnect) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DhConnect) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DhConnect) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *DhConnect) List(filters map[string]interface{}) []*DhConnect {
	var list []*DhConnect
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhConnect) OrderList(filters map[string]interface{},order ...string) []*DhConnect {
	var list []*DhConnect
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhConnect) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhConnect) {
	var list []*DhConnect
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}