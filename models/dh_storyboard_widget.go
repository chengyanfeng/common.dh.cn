package models

import (
	"time"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type DhStoryboardWidget struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    PageId string `json:"page_id"`
    Type string `json:"type"`
    Position string `json:"position"`
    Content string `json:"content"`
    Status int `json:"status"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhStoryboardWidget))
}

func (m *DhStoryboardWidget) TableName() string {
    return "dh_storyboard_widget"
}

func (m *DhStoryboardWidget) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhStoryboardWidget) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhStoryboardWidget) Find(args ...interface{}) *DhStoryboardWidget {
	data := m.find(m,args...)
	if data != nil {
		return data.(*DhStoryboardWidget)
	} else {
		return nil
	}
}

func (m *DhStoryboardWidget) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *DhStoryboardWidget) List(filters map[string]interface{}) []*DhStoryboardWidget {
	var list []*DhStoryboardWidget
	_, err := m.findByFilters(m,filters).All(&list)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return list
}

func (m *DhStoryboardWidget) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhStoryboardWidget) {
	var list []*DhStoryboardWidget
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m,page, page_size, filters).All(&list)
	if err != nil {
		utils.Error(err)
		return 0,0,nil
	}
	return total, total_page, list
}