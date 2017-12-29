package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DhStoryboardPage struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    	BoardId string `json:"board_id"`
    	Thumbnail string `json:"thumbnail"`
    	Status int `json:"status"`
    	Sort int `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhStoryboardPage))
}

func (m *DhStoryboardPage) TableName() string {
    return "dh_storyboard_page"
}

func (m *DhStoryboardPage) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhStoryboardPage) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhStoryboardPage) Find(args ...interface{}) *DhStoryboardPage {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DhStoryboardPage)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhStoryboardPage) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DhStoryboardPage) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DhStoryboardPage) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *DhStoryboardPage) List(filters map[string]interface{}) []*DhStoryboardPage {
	var list []*DhStoryboardPage
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhStoryboardPage) OrderList(filters map[string]interface{}, order ...string) []*DhStoryboardPage {
	var list []*DhStoryboardPage
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhStoryboardPage) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhStoryboardPage) {
	var list []*DhStoryboardPage
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}

func (m *DhStoryboardPage) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DhStoryboardPage) {
	var list []*DhStoryboardPage
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}