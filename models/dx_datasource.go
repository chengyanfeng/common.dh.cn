package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type DxDatasource struct {
	DhBase
	Id              int64     `json:"-"`
	ObjectId        string    `json:"_id"`
	Name            string    `json:"name"`
	Type            string    `json:"type"`
	Format          string    `json:"format"`
	Url             string    `json:"url"`
	ConnectId       string    `json:"connect_id"`
	Table           string    `json:"table"`
	Sql             string    `json:"sql"`
	Data            string    `json:"data"`
	IsAutoUpdate    int       `json:"is_auto_update"`
	UpdateFrequency int       `json:"update_frequency"`
	Spec            string    `json:"spec"`
	Sort            int       `json:"sort"`
	Status          int       `json:"status"`
	CreateTime      time.Time `json:"-"`
	UpdateTime      time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DxDatasource))
}

func (m *DxDatasource) TableName() string {
	return "dx_datasource"
}

func (m *DxDatasource) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DxDatasource) Save() bool {
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DxDatasource) Find(args ...interface{}) *DxDatasource {
	data := m.find(m, args...)
	if data != nil {
		_data, ok := data.(*DxDatasource)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DxDatasource) Delete(args ...interface{}) bool {
	return m.delete(m, args...)
}

func (m *DxDatasource) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m, args...)
}

func (m *DxDatasource) Count(filters map[string]interface{}) int64 {
	return m.count(m, filters)
}

func (m *DxDatasource) List(filters map[string]interface{}) []*DxDatasource {
	var list []*DxDatasource
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DxDatasource) OrderList(filters map[string]interface{}, order ...string) []*DxDatasource {
	var list []*DxDatasource
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DxDatasource) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DxDatasource) {
	var list []*DxDatasource
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DxDatasource) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DxDatasource) {
	var list []*DxDatasource
	total, total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0, 0, nil
	}
	return total, total_page, list
}

func (m *DxDatasource) SortById(_ids []string) bool {
	o := new(DhBase).Orm()
	for i, _id := range _ids {
		ds := new(DxDatasource).Find("object_id", _id)
		if ds != nil {
			ds.Sort = i
			if ok := ds.Save(); !ok {
				o.Rollback()
				return false
			}
		}
	}
	o.Commit()

	return true
}
