package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhRelation struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
	CorpId string `json:"corp_id"`
	UserId string `json:"user_id"`
	RelateType string `json:"relate_type"`
	RelateId string `json:"relate_id"`
	Auth string `json:"auth"`
	Name string `json:"name"`
	Sort int `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhRelation))
}

func (m *DhRelation) TableName() string {
	return "dh_relation"
}

func (m *DhRelation) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhRelation) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhRelation) Find(args ...interface{}) *DhRelation {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhRelation)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhRelation) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DhRelation) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DhRelation) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *DhRelation) List(filters map[string]interface{}) []*DhRelation {
	var list []*DhRelation
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhRelation) OrderList(filters map[string]interface{}, order ...string) []*DhRelation {
	var list []*DhRelation
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhRelation) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhRelation) {
	var list []*DhRelation
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}

func (m *DhRelation) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DhRelation) {
	var list []*DhRelation
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}