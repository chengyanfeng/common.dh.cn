package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type {{.ModelName}} struct {
	DhBase
	Id int64 `json:"-"`
	ObjectId string `json:"_id"`
	{{- range .Fields}}
    {{.Name}} {{.Type}} {{.Tag}}
	{{- end}}
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
		_data,ok := data.(*{{.ModelName}})
		if ok {
			return _data
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (m *{{.ModelName}}) Delete(args ...interface{}) bool {
	return m.delete(m,args...)
}

func (m *{{.ModelName}}) SoftDelete(args ...interface{}) bool {
	return m.softDelete(m,args...)
}

func (m *{{.ModelName}}) Count(filters map[string]interface{}) int64 {
	return m.count(m,filters)
}

func (m *{{.ModelName}}) List(filters map[string]interface{}) []*{{.ModelName}} {
	var list []*{{.ModelName}}
	_, err := m.findByFilters(m, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *{{.ModelName}}) OrderList(filters map[string]interface{}, order ...string) []*{{.ModelName}} {
	var list []*{{.ModelName}}
	_, err := m.findByFilters(m, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return nil
	}
	return list
}

func (m *{{.ModelName}}) Pager(page int64, page_size int64, filters map[string]interface{}) (total int64, total_page int64, result []*{{.ModelName}}) {
	var list []*{{.ModelName}}
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}

func (m *{{.ModelName}}) OrderPager(page int64, page_size int64, filters map[string]interface{}, order ...string) (total int64, total_page int64, result []*{{.ModelName}}) {
	var list []*{{.ModelName}}
	total,total_page = m.pager(m, filters, page_size)
	_, err := m.pagerList(m, page, page_size, filters).OrderBy(order...).All(&list)
	if err != nil {
		m.errReport(err)
		return 0,0,nil
	}
	return total, total_page, list
}