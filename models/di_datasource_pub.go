package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiDatasourcePub struct {
	DiBase
	Id               int64     `json:"-"`
	ObjectId         string    `json:"_id"`
	Name             string    `json:"name"`
	Status           int       `json:"status"`
	DatasourceTypeId string    `json:"datasource_type_id"`
	Logo             string    `json:"logo"`
	DiDatasourceId   string    `json:"di_datasource_id"`
	CreateTime       time.Time `json:"-"`
	UpdateTime       time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiDatasourcePub))
}

func (m *DiDatasourcePub) TableName() string {
	return "di_datasource_pub"
}

func (m *DiDatasourcePub) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiDatasourcePub) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiDatasourcePub) Find(args ...interface{}) *DiDatasourcePub {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiDatasourcePub)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiDatasourcePub) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiDatasourcePub) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiDatasourcePub) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiDatasourcePub) List(filters map[string]interface{}) []*DiDatasourcePub {
	var list []*DiDatasourcePub
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDatasourcePub) OrderList(filters map[string]interface{}, order ...string) []*DiDatasourcePub {
	var list []*DiDatasourcePub
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDatasourcePub) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiDatasourcePub) {
	var list []*DiDatasourcePub
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiDatasourcePub) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiDatasourcePub) {
	var list []*DiDatasourcePub
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
