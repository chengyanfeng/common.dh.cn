package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)
type DhList struct {
	DhBase
	Id          int64     `json:"-"`
	ObjectId    string    `json:"_id"`
	Path        string    `json:"path"`
	Name        string    `json:"name"`
	FObjectId   string `json:"f_object_id"`
	CreateTime  time.Time `json:"-"`
	UpdateTime  time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhList))
}
func (m *DhList) TableName() string {
	return "dh_list"
}

func (m *DhList) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhList) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhList) Find(args ...interface{}) *DhList {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DhList)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhList) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DhList) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DhList) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DhList) List(filters map[string]interface{}) []*DhList {
	var list []*DhList
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhList) OrderList(filters map[string]interface{}, order ...string) []*DhList {
	var list []*DhList
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhList) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhList) {
	var list []*DhList
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DhList) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DhList) {
	var list []*DhList
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
