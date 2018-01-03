package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DxScreenWidget struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
	ScreenId string `json:"screen_id"`
	DatasourceId string `json:"datasource_id"`
	Config  `json:"config"`
	Status int `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DxScreenWidget))
}

func (m *DxScreenWidget) TableName() string {
	return "dx_screen_widget"
}

func (m *DxScreenWidget) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DxScreenWidget) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DxScreenWidget) Find(args ...interface{}) *DxScreenWidget {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DxScreenWidget)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DxScreenWidget) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DxScreenWidget) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DxScreenWidget) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *DxScreenWidget) List(filters map[string]interface{}) []*DxScreenWidget {
	var list []*DxScreenWidget
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DxScreenWidget) OrderList(filters map[string]interface{}, order ...string) []*DxScreenWidget {
	var list []*DxScreenWidget
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DxScreenWidget) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DxScreenWidget) {
	var list []*DxScreenWidget
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}

func (m *DxScreenWidget) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DxScreenWidget) {
	var list []*DxScreenWidget
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}