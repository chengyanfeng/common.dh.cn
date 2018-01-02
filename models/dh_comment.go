package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhComment struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    RelateType string `json:"relate_type"`
    RelateId string `json:"relate_id"`
    UserId string `json:"user_id"`
    Type string `json:"type"`
    Content string `json:"content"`
    ReplayId string `json:"replay_id"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhComment))
}

func (m *DhComment) TableName() string {
    return "dh_comment"
}

func (m *DhComment) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhComment) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhComment) Find(args ...interface{}) *DhComment {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhComment)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhComment) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DhComment) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DhComment) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *DhComment) List(filters map[string]interface{}) []*DhComment {
	var list []*DhComment
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhComment) OrderList(filters map[string]interface{}, order ...string) []*DhComment {
	var list []*DhComment
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhComment) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhComment) {
	var list []*DhComment
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}

func (m *DhComment) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DhComment) {
	var list []*DhComment
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}