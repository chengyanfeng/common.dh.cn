package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type DxScreen struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
	Name string `json:"name"`
	Config string `json:"config"`
	Thumbnail string `json:"thumbnail"`
	Status int `json:"status"`
	Sort int `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DxScreen))
}

func (m *DxScreen) TableName() string {
	return "dx_screen"
}

func (m *DxScreen) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DxScreen) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DxScreen) Find(args ...interface{}) *DxScreen {
	data := m.find(m,args...)
	if data != nil {
		_data,ok := data.(*DxScreen)
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *DxScreen) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *DxScreen) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *DxScreen) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *DxScreen) List(filters map[string]interface{}) []*DxScreen {
	var list []*DxScreen
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DxScreen) OrderList(filters map[string]interface{}, order ...string) []*DxScreen {
	var list []*DxScreen
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *DxScreen) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DxScreen) {
	var list []*DxScreen
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}

func (m *DxScreen) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*DxScreen) {
	var list []*DxScreen
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}