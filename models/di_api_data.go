package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiApiData struct {
	DiBase
	Id         int64     `json:"-"`
	ObjectId   string    `json:"_id"`
	UserId     string    `json:"user_id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiApiData))
}

func (m *DiApiData) TableName() string {
	return "di_api_data"
}

func (m *DiApiData) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiApiData) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiApiData) Find(args ...interface{}) *DiApiData {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiApiData)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiApiData) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiApiData) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiApiData) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiApiData) List(filters map[string]interface{}) []*DiApiData {
	var list []*DiApiData
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiApiData) OrderList(filters map[string]interface{}, order ...string) []*DiApiData {
	var list []*DiApiData
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiApiData) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiApiData) {
	var list []*DiApiData
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiApiData) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiApiData) {
	var list []*DiApiData
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
