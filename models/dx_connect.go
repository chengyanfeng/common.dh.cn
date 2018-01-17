package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DxConnect struct {
	DxBase
	Id         int64     `json:"-"`
	ObjectId   string    `json:"_id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Config     string    `json:"config"`
	ErrMsg     string    `json:"err_msg"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DxConnect))
}

func (m *DxConnect) TableName() string {
	return "dx_connect"
}

func (m *DxConnect) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DxConnect) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DxConnect) Find(args ...interface{}) *DxConnect {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DxConnect)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DxConnect) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DxConnect) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DxConnect) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DxConnect) List(filters map[string]interface{}) []*DxConnect {
	var list []*DxConnect
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DxConnect) OrderList(filters map[string]interface{}, order ...string) []*DxConnect {
	var list []*DxConnect
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DxConnect) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DxConnect) {
	var list []*DxConnect
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DxConnect) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DxConnect) {
	var list []*DxConnect
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
