package models

import (
	"time"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type DhStoryboardPage struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    BoardId string `json:"board_id"`
    Thumbnail string `json:"thumbnail"`
    Status int `json:"status"`
    Sort int `json:"sort"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhStoryboardPage))
}

func (m *DhStoryboardPage) TableName() string {
    return "dh_storyboard_page"
}

func (m *DhStoryboardPage) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhStoryboardPage) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhStoryboardPage) Find(args ...interface{}) *DhStoryboardPage {
	data := m.find(m,args...)
	if data != nil {
		return data.(*DhStoryboardPage)
	} else {
		return nil
	}
}

func (m *DhStoryboardPage) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *DhStoryboardPage) List(filters map[string]interface{}) []*DhStoryboardPage {
	var list []*DhStoryboardPage
	_, err := m.findByFilters(m,filters).All(&list)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return list
}

func (m *DhStoryboardPage) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhStoryboardPage) {
	var list []*DhStoryboardPage
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m,page, page_size, filters).All(&list)
	if err != nil {
		utils.Error(err)
		return 0,0,nil
	}
	return total, total_page, list
}