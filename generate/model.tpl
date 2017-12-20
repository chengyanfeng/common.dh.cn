package models

import (
	"time"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type {{.ModelName}} struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`{{range .Fields}}
    {{.Name}} {{.Type}} {{.Tag}}{{end}}
	CreateTime time.Time `json:"-"`
	UpdateTime time.Time `json:"-"`
}

func init() {
	orm.RegisterModel(new({{.ModelName}}))
}

func (m *{{.ModelName}}) TableName() string {
    return "{{.TableName}}"
}

func (m *{{.ModelName}}) Query() orm.QuerySeter {
	return m.query(m)
}

func (m *{{.ModelName}}) Save() bool{
	if m.Id == 0 {
		return m.create(m)
	} else {
		return m.update(m)
	}
}

func (m *{{.ModelName}}) Find(args ...interface{}) *{{.ModelName}} {
	data := m.find(m,args...)
	if data != nil {
		return data.(*{{.ModelName}})
	} else {
		return nil
	}
}

func (m *{{.ModelName}}) Delete(index interface{}) bool {
	return m.delete(m,index)
}

func (m *{{.ModelName}}) List(filters map[string]interface{}) []*{{.ModelName}} {
	var list []*{{.ModelName}}
	_, err := m.findByFilters(m,filters).All(&list)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return list
}

func (m *{{.ModelName}}) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*{{.ModelName}}) {
	var list []*{{.ModelName}}
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m,page, page_size, filters).All(&list)
	if err != nil {
		utils.Error(err)
		return 0,0,nil
	}
	return total, total_page, list
}