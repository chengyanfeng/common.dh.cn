package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhVcode struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    	Mobile string `json:"mobile"`
    	Code string `json:"code"`
    	Status int `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhVcode))
}

func (m *DhVcode) TableName() string {
    return "dh_vcode"
}

func (m *DhVcode) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhVcode) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhVcode) Find(args ...interface{}) *DhVcode {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhVcode)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhVcode) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DhVcode) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DhVcode) List(filters map[string]interface{}) []*DhVcode {
	var list []*DhVcode
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhVcode) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhVcode) {
	var list []*DhVcode
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}