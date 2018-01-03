package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhStoryboardGroup struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
	Name string `json:"name"`
	Status int `json:"status"`
	Sort int `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhStoryboardGroup))
}

func (m *DhStoryboardGroup) TableName() string {
	return "dh_storyboard_group"
}

func (m *DhStoryboardGroup) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhStoryboardGroup) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhStoryboardGroup) Find(args ...interface{}) *DhStoryboardGroup {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhStoryboardGroup)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhStoryboardGroup) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DhStoryboardGroup) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DhStoryboardGroup) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *DhStoryboardGroup) List(filters map[string]interface{}) []*DhStoryboardGroup {
	var list []*DhStoryboardGroup
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhStoryboardGroup) OrderList(filters map[string]interface{}, order ...string) []*DhStoryboardGroup {
	var list []*DhStoryboardGroup
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhStoryboardGroup) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhStoryboardGroup) {
	var list []*DhStoryboardGroup
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}

func (m *DhStoryboardGroup) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DhStoryboardGroup) {
	var list []*DhStoryboardGroup
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}