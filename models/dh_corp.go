package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DhCorp struct {
	DhBase
	Id         int64     `json:"-"`
	ObjectId   string    `json:"_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Mobile     string    `json:"mobile"`
	Vcode      string    `json:"vcode"`
	ConnectId  string    `json:"connect_id"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhCorp))
}

func (m *DhCorp) TableName() string {
	return "dh_corp"
}

func (m *DhCorp) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhCorp) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhCorp) Find(args ...interface{}) *DhCorp {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DhCorp)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DhCorp) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DhCorp) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DhCorp) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DhCorp) List(filters map[string]interface{}) []*DhCorp {
	var list []*DhCorp
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhCorp) OrderList(filters map[string]interface{}, order ...string) []*DhCorp {
	var list []*DhCorp
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DhCorp) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhCorp) {
	var list []*DhCorp
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DhCorp) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DhCorp) {
	var list []*DhCorp
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
