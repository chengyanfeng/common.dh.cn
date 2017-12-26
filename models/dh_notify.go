package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhNotify struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    	FromCropId string `json:"from_crop_id"`
    	FromUserId string `json:"from_user_id"`
    	UserId string `json:"user_id"`
    	Type string `json:"type"`
    	Config string `json:"config"`
    	Status int `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhNotify))
}

func (m *DhNotify) TableName() string {
    return "dh_notify"
}

func (m *DhNotify) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhNotify) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhNotify) Find(args ...interface{}) *DhNotify {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhNotify)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhNotify) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DhNotify) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DhNotify) List(filters map[string]interface{}) []*DhNotify {
	var list []*DhNotify
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhNotify) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhNotify) {
	var list []*DhNotify
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}