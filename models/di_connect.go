package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiConnect struct {
	DiBase
	Id         int64     `json:"-"`
	ObjectId   string    `json:"_id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Format     string    `json:"format"`
	Config     string    `json:"config"`
	ErrMsg     string    `json:"err_msg"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiConnect))
}

func (m *DiConnect) TableName() string {
	return "di_connect"
}

func (m *DiConnect) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiConnect) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiConnect) Find(args ...interface{}) *DiConnect {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiConnect)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiConnect) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiConnect) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiConnect) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiConnect) List(filters map[string]interface{}) []*DiConnect {
	var list []*DiConnect
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiConnect) OrderList(filters map[string]interface{}, order ...string) []*DiConnect {
	var list []*DiConnect
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiConnect) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiConnect) {
	var list []*DiConnect
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiConnect) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiConnect) {
	var list []*DiConnect
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
