package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiStoryboardGroup struct {
	DiBase
	Id         int64     `json:"-"`
	ObjectId   string    `json:"_id"`
	Name       string    `json:"name"`
	Status     int       `json:"status"`
	Sort       int       `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiStoryboardGroup))
}

func (m *DiStoryboardGroup) TableName() string {
	return "di_storyboard_group"
}

func (m *DiStoryboardGroup) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiStoryboardGroup) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiStoryboardGroup) Find(args ...interface{}) *DiStoryboardGroup {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiStoryboardGroup)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiStoryboardGroup) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiStoryboardGroup) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiStoryboardGroup) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiStoryboardGroup) List(filters map[string]interface{}) []*DiStoryboardGroup {
	var list []*DiStoryboardGroup
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiStoryboardGroup) OrderList(filters map[string]interface{}, order ...string) []*DiStoryboardGroup {
	var list []*DiStoryboardGroup
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiStoryboardGroup) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiStoryboardGroup) {
	var list []*DiStoryboardGroup
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiStoryboardGroup) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiStoryboardGroup) {
	var list []*DiStoryboardGroup
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
