package models

import (
	"time"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type DhUserCorp struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
    UserId string `json:"user_id"`
    CropId string `json:"crop_id"`
    Role string `json:"role"`
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new(DhUserCorp))
}

func (m *DhUserCorp) TableName() string {
    return "dh_user_corp"
}

func (m *DhUserCorp) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *DhUserCorp) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *DhUserCorp) Find(args ...interface{}) *DhUserCorp {
	data := m.find(m,args...)
	if data != nil {
		return data.(*DhUserCorp)
	} else {
		return nil
	}
}

func (m *DhUserCorp) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *DhUserCorp) List(filters map[string]interface{}) []*DhUserCorp {
	var list []*DhUserCorp
	_, err := m.findByFilters(m,filters).All(&list)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return list
}

func (m *DhUserCorp) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*DhUserCorp) {
	var list []*DhUserCorp
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m,page, page_size, filters).All(&list)
	if err != nil {
		utils.Error(err)
		return 0,0,nil
	}
	return total, total_page, list
}