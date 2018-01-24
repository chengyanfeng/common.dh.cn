package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiLinkshare struct {
	DiBase
	Id         int64     `json:"-"`
	ObjectId   string    `json:"_id"`
	RelateType string    `json:"relate_type"`
	RelateId   string    `json:"relate_id"`
	Type       int       `json:"type"`
	Password   string    `json:"password"`
	Url        string    `json:"url"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiLinkshare))
}

func (m *DiLinkshare) TableName() string {
	return "di_linkshare"
}

func (m *DiLinkshare) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiLinkshare) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiLinkshare) Find(args ...interface{}) *DiLinkshare {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiLinkshare)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiLinkshare) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiLinkshare) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiLinkshare) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiLinkshare) List(filters map[string]interface{}) []*DiLinkshare {
	var list []*DiLinkshare
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiLinkshare) OrderList(filters map[string]interface{}, order ...string) []*DiLinkshare {
	var list []*DiLinkshare
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiLinkshare) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiLinkshare) {
	var list []*DiLinkshare
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiLinkshare) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiLinkshare) {
	var list []*DiLinkshare
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
