package models

import (
	"time"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type DhStoryboard struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    	GroupId string `json:"group_id"`
    	Name string `json:"name"`
    	Status int `json:"status"`
    	Sort int `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhStoryboard))
}

func (m *DhStoryboard) TableName() string {
    return "dh_storyboard"
}

func (m *DhStoryboard) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhStoryboard) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhStoryboard) Find(args ...interface{}) *DhStoryboard {
	data := m.find(m,args...)
	if data != nil {
		return data.(*DhStoryboard)
	} else {
		return nil
	}
}

func (m *DhStoryboard) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *DhStoryboard) List(filters map[string]interface{}) []*DhStoryboard {
	var list []*DhStoryboard
	_, err := m.findByFilters(m,filters).All(&list)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return list
}

func (m *DhStoryboard) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhStoryboard) {
	var list []*DhStoryboard
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m,page, page_size, filters).All(&list)
	if err != nil {
		utils.Error(err)
		return 0,0,nil
	}
	return total, total_page, list
}