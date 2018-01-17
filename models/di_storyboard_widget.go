package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiStoryboardWidget struct {
	DiBase
	Id         int64     `json:"-"`
	ObjectId   string    `json:"_id"`
	PageId     string    `json:"page_id"`
	Type       string    `json:"type"`
	Position   string    `json:"position"`
	Content    string    `json:"content"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiStoryboardWidget))
}

func (m *DiStoryboardWidget) TableName() string {
	return "di_storyboard_widget"
}

func (m *DiStoryboardWidget) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiStoryboardWidget) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiStoryboardWidget) Find(args ...interface{}) *DiStoryboardWidget {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiStoryboardWidget)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiStoryboardWidget) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiStoryboardWidget) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiStoryboardWidget) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiStoryboardWidget) List(filters map[string]interface{}) []*DiStoryboardWidget {
	var list []*DiStoryboardWidget
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiStoryboardWidget) OrderList(filters map[string]interface{}, order ...string) []*DiStoryboardWidget {
	var list []*DiStoryboardWidget
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiStoryboardWidget) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiStoryboardWidget) {
	var list []*DiStoryboardWidget
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiStoryboardWidget) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiStoryboardWidget) {
	var list []*DiStoryboardWidget
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
