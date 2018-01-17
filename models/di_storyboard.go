package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiStoryboard struct {
	DiBase
	Id         int64     `json:"-"`
	ObjectId   string    `json:"_id"`
	GroupId    string    `json:"group_id"`
	Name       string    `json:"name"`
	Status     int       `json:"status"`
	Sort       int       `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiStoryboard))
}

func (m *DiStoryboard) TableName() string {
	return "di_storyboard"
}

func (m *DiStoryboard) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiStoryboard) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiStoryboard) Find(args ...interface{}) *DiStoryboard {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiStoryboard)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiStoryboard) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiStoryboard) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiStoryboard) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiStoryboard) List(filters map[string]interface{}) []*DiStoryboard {
	var list []*DiStoryboard
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiStoryboard) OrderList(filters map[string]interface{}, order ...string) []*DiStoryboard {
	var list []*DiStoryboard
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiStoryboard) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiStoryboard) {
	var list []*DiStoryboard
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiStoryboard) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiStoryboard) {
	var list []*DiStoryboard
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
