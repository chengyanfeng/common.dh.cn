package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiNotify struct {
	DiBase
	Id         int64     `json:"-"`
	ObjectId   string    `json:"_id"`
	FromCropId string    `json:"from_crop_id"`
	FromUserId string    `json:"from_user_id"`
	UserId     string    `json:"user_id"`
	Type       string    `json:"type"`
	Config     string    `json:"config"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiNotify))
}

func (m *DiNotify) TableName() string {
	return "di_notify"
}

func (m *DiNotify) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiNotify) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiNotify) Find(args ...interface{}) *DiNotify {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiNotify)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiNotify) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiNotify) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiNotify) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiNotify) List(filters map[string]interface{}) []*DiNotify {
	var list []*DiNotify
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiNotify) OrderList(filters map[string]interface{}, order ...string) []*DiNotify {
	var list []*DiNotify
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiNotify) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiNotify) {
	var list []*DiNotify
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiNotify) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiNotify) {
	var list []*DiNotify
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
