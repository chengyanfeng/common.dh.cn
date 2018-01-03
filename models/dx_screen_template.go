package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DxScreenTemplate struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
	ScreenId string `json:"screen_id"`
	Status int `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DxScreenTemplate))
}

func (m *DxScreenTemplate) TableName() string {
	return "dx_screen_template"
}

func (m *DxScreenTemplate) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DxScreenTemplate) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DxScreenTemplate) Find(args ...interface{}) *DxScreenTemplate {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DxScreenTemplate)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DxScreenTemplate) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DxScreenTemplate) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DxScreenTemplate) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *DxScreenTemplate) List(filters map[string]interface{}) []*DxScreenTemplate {
	var list []*DxScreenTemplate
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DxScreenTemplate) OrderList(filters map[string]interface{}, order ...string) []*DxScreenTemplate {
	var list []*DxScreenTemplate
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DxScreenTemplate) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DxScreenTemplate) {
	var list []*DxScreenTemplate
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}

func (m *DxScreenTemplate) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DxScreenTemplate) {
	var list []*DxScreenTemplate
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}