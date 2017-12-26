package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhIcode struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    	Code string `json:"code"`
    	Status int `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhIcode))
}

func (m *DhIcode) TableName() string {
    return "dh_icode"
}

func (m *DhIcode) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhIcode) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhIcode) Find(args ...interface{}) *DhIcode {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhIcode)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhIcode) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DhIcode) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DhIcode) List(filters map[string]interface{}) []*DhIcode {
	var list []*DhIcode
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhIcode) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhIcode) {
	var list []*DhIcode
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}