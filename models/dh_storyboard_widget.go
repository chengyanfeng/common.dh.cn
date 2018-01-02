package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhStoryboardWidget struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
	PageId string `json:"page_id"`
	Type string `json:"type"`
	Position string `json:"position"`
	Content string `json:"content"`
	Status int `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhStoryboardWidget))
}

func (m *DhStoryboardWidget) TableName() string {
    return "dh_storyboard_widget"
}

func (m *DhStoryboardWidget) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhStoryboardWidget) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhStoryboardWidget) Find(args ...interface{}) *DhStoryboardWidget {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhStoryboardWidget)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhStoryboardWidget) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DhStoryboardWidget) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DhStoryboardWidget) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *DhStoryboardWidget) List(filters map[string]interface{}) []*DhStoryboardWidget {
	var list []*DhStoryboardWidget
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhStoryboardWidget) OrderList(filters map[string]interface{}, order ...string) []*DhStoryboardWidget {
	var list []*DhStoryboardWidget
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhStoryboardWidget) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhStoryboardWidget) {
	var list []*DhStoryboardWidget
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}

func (m *DhStoryboardWidget) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DhStoryboardWidget) {
	var list []*DhStoryboardWidget
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}