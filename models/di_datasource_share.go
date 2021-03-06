package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiDataSourceShare struct {
	DiBase
	Id               int64     `json:"-"`
	ObjectId         string    `json:"_id"`
	Name             string    `json:"name"`
	Status           int       `json:"status"`
	UserId 			string    `json:"user_id"`
	CorpId            string    `json:"corp_id"`
	DatasourceId   string    `json:"datasource_id"`
	Fields       string    `json:"fields"`
	IsFullShow       string    `json:"is_full_show"`

	CreateTime       time.Time `json:"-"`
	UpdateTime       time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiDataSourceShare))
}

func (m *DiDataSourceShare) TableName() string {
	return "di_datasource_share"
}

func (m *DiDataSourceShare) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiDataSourceShare) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiDataSourceShare) Find(args ...interface{}) *DiDataSourceShare {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiDataSourceShare)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiDataSourceShare) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiDataSourceShare) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiDataSourceShare) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiDataSourceShare) List(filters map[string]interface{}) []*DiDataSourceShare {
	var list []*DiDataSourceShare
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDataSourceShare) OrderList(filters map[string]interface{}, order ...string) []*DiDataSourceShare {
	var list []*DiDataSourceShare
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDataSourceShare) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiDataSourceShare) {
	var list []*DiDataSourceShare
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiDataSourceShare) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiDataSourceShare) {
	var list []*DiDataSourceShare
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
