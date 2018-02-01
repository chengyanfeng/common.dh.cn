package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DiDatasourceModify struct {
	DiBase
	Id           int64     `json:"-"`
	ObjectId     string    `json:"_id"`
	UserId       string    `json:"user_id"`
	DatasourceId string    `json:"datasource_id"`
	Table        string    `json:"table"`
	RowId        string    `json:"row_id"`
	Column       string    `json:"column"`
	Type         string    `json:"type"`
	OldValue     string    `json:"old_value"`
	NewValue     string    `json:"new_value"`
	CreateTime   time.Time `json:"-"`
	UpdateTime   time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DiDatasourceModify))
}

func (m *DiDatasourceModify) TableName() string {
	return "di_datasource_modify"
}

func (m *DiDatasourceModify) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DiDatasourceModify) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DiDatasourceModify) Find(args ...interface{}) *DiDatasourceModify {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DiDatasourceModify)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DiDatasourceModify) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DiDatasourceModify) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DiDatasourceModify) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DiDatasourceModify) List(filters map[string]interface{}) []*DiDatasourceModify {
	var list []*DiDatasourceModify
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDatasourceModify) OrderList(filters map[string]interface{}, order ...string) []*DiDatasourceModify {
	var list []*DiDatasourceModify
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DiDatasourceModify) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DiDatasourceModify) {
	var list []*DiDatasourceModify
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DiDatasourceModify) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DiDatasourceModify) {
	var list []*DiDatasourceModify
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}
