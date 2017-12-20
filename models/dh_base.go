package models

import (
	"time"
	"fmt"
	"reflect"
	"github.com/rs/xid"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"common.dh.cn/utils"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	host := beego.AppConfig.String("mysql_host")
	port := beego.AppConfig.String("mysql_port")
	name := beego.AppConfig.String("mysql_name")
	username := beego.AppConfig.String("mysql_username")
	password := beego.AppConfig.String("mysql_password")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, name)
	orm.RegisterDataBase("default", "mysql", connection)
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.DefaultTimeLoc = time.UTC
	orm.Debug = true
}

type DhBase struct {
}

func (m *DhBase) Orm() orm.Ormer {
	return orm.NewOrm()
}

func (m *DhBase) query(entity interface{}) orm.QuerySeter {
	return m.Orm().QueryTable(entity)
}

func (m *DhBase) create(entity interface{}) bool {
	mutable := reflect.ValueOf(entity).Elem()
	mutable.FieldByName("ObjectId").SetString(xid.New().String())
	now := time.Now();
	mutable.FieldByName("CreateTime").Set(reflect.ValueOf(now))
	mutable.FieldByName("UpdateTime").Set(reflect.ValueOf(now))
	_id, err := orm.NewOrm().Insert(entity)
	if err != nil {
		return false
	} else {
		mutable.FieldByName("Id").SetInt(_id)
		return true
	}
}

func (m *DhBase) update(entity interface{}) bool {
	mutable := reflect.ValueOf(entity).Elem()
	now := time.Now();
	mutable.FieldByName("UpdateTime").Set(reflect.ValueOf(now))
	_, err := orm.NewOrm().Update(entity)
	if err != nil {
		utils.Error(err)
		return false
	} else {
		return true
	}
}

func (m *DhBase) delete(entity interface{}, index interface{}) bool {
	var err interface{}
	switch index.(type) {
		case string,*string:
			_, err = m.query(entity).Filter("object_id", index).Delete()
		case int,*int:
			_, err = m.query(entity).Filter("id", index).Delete()
		default:
	}
	if err != nil {
		utils.Error(err)
		return false
	}
	return true
}

func (m *DhBase) find(entity interface{},args ...interface{}) interface{} {
	total := len(args)
	if total == 1 {
		index := args[0]
		switch index.(type) {
			case string,*string:
				return m.findByObjectID(entity,index.(string))
			case int,*int:
				return m.findByID(entity,index.(int64))
			default:
				return nil
		}
	} else if total == 2 {
		key := args[0]
		value := args[0]
		_key,ok := key.(string)
		if !ok {
			return nil
		}
		return m.findByFilter(entity, _key, value)
	} else {
		return nil
	}
}

func (m *DhBase) findByFilter(entity interface{}, key string, value interface {}) (result interface{}) {
	err := m.query(entity).Filter(key, value).One(entity)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return entity
}

func (m *DhBase) findByFilters(entity interface{}, filters map[string]interface{}) (result interface{}) {
	query := m.query(entity)
	for k,v := range filters {
		query = query.Filter(k,v)
	}
	err := query.One(entity)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return entity
}

func (m *DhBase) findByID(entity interface{}, id int64) (result interface{}) {
	err := m.query(m).Filter("id", id).One(entity)
	if err != nil {
		utils.Error(err)
		return nil
	}
	return entity
}

func (m *DhBase) findByObjectID(entity interface{}, object_id string) (result interface{}) {
	err := m.query(entity).Filter("object_id", object_id).One(entity)
	if err != nil {
		if err.Error() != "<QuerySeter> no row found" {
			utils.Error(err)
		}
		return nil
	}
	return entity
}

func (m *DhBase) filter(entity interface{}, filters map[string]interface{}) orm.QuerySeter {
	query := m.query(entity)
	for k,v := range filters {
		query = query.Filter(k,v)
	}
	return query
}

func (m *DhBase) count(entity interface{}, filters map[string]interface{}) int64 {
	result, err := m.filter(m,filters).Count()
	if err != nil {
		utils.Error(err)
		return 0
	}
	return result
}

func (m *DhBase) pager(entity interface{}, page int64, page_size int64, filters map[string]interface{}) orm.QuerySeter {
	return m.filter(m,filters).Offset(page * page_size).Limit(page_size)
}