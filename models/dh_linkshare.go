package models

import (
	"time"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type DhLinkshare struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    	RelateType string `json:"relate_type"`
    	RelateId string `json:"relate_id"`
    	Type string `json:"type"`
    	Password string `json:"password"`
    	Url string `json:"url"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhLinkshare))
}

func (m *DhLinkshare) TableName() string {
    return "dh_linkshare"
}

func (m *DhLinkshare) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhLinkshare) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhLinkshare) Find(args ...interface{}) *DhLinkshare {
	data := m.find(m,args...)
	if data != nil {
		return data.(*DhLinkshare)
	} else {
		return nil
	}
}

func (m *DhLinkshare) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *DhLinkshare) List(filters map[string]interface{}) []*DhLinkshare {
	var list []*DhLinkshare
	_, err := m.findByFilters(m,filters).All(&list)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return list
}

func (m *DhLinkshare) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhLinkshare) {
	var list []*DhLinkshare
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m,page, page_size, filters).All(&list)
	if err != nil {
		utils.Error(err)
		return 0,0,nil
	}
	return total, total_page, list
}