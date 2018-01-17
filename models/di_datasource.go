package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiDatasource struct {
	DiBase
	Id              int64     `json:"-"`
	ObjectId        string    `json:"_id"`
	GroupId         string    `json:"group_id"`
	Name            string    `json:"name"`
	Table           string    `json:"table"`
	Type            string    `json:"type"`
	Format          string    `json:"format"`
	Url             string    `json:"url"`
	ConnectId       string    `json:"connect_id"`
	IsAutoUpdate    int       `json:"is_auto_update"`
	UpdateFrequency int       `json:"update_frequency"`
	UpdateLocation  int64     `json:"update_location"`
	Sort            int       `json:"sort"`
	Status          int       `json:"status"`
	CreateTime      time.Time `json:"-"`
	UpdateTime      time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiDatasource))
}

func (m *DiDatasource) TableName() string {
	return "di_datasource"
}

func (m *DiDatasource) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiDatasource) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiDatasource) Find(args ...interface{}) *DiDatasource {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiDatasource)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiDatasource) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiDatasource) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiDatasource) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiDatasource) List(filters map[string]interface{}) []*DiDatasource {
	var list []*DiDatasource
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDatasource) OrderList(filters map[string]interface{}, order ...string) []*DiDatasource {
	var list []*DiDatasource
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDatasource) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiDatasource) {
	var list []*DiDatasource
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiDatasource) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiDatasource) {
	var list []*DiDatasource
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
