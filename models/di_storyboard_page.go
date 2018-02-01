package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiStoryboardPage struct {
	DiBase
	Id         int64     `json:"-"`
	ObjectId   string    `json:"_id"`
	BoardId    string    `json:"board_id"`
	Thumbnail  string    `json:"thumbnail"`
	Status     int       `json:"status"`
	Sort       int       `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiStoryboardPage))
}

func (m *DiStoryboardPage) TableName() string {
	return "di_storyboard_page"
}

func (m *DiStoryboardPage) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiStoryboardPage) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiStoryboardPage) Find(args ...interface{}) *DiStoryboardPage {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiStoryboardPage)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiStoryboardPage) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiStoryboardPage) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiStoryboardPage) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiStoryboardPage) List(filters map[string]interface{}) []*DiStoryboardPage {
	var list []*DiStoryboardPage
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiStoryboardPage) OrderList(filters map[string]interface{}, order ...string) []*DiStoryboardPage {
	var list []*DiStoryboardPage
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiStoryboardPage) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiStoryboardPage) {
	var list []*DiStoryboardPage
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiStoryboardPage) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiStoryboardPage) {
	var list []*DiStoryboardPage
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
