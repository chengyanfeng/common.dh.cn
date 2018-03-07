package models

import (
	"fmt"
	"math"
	"net/url"
	"reflect"
	"strings"

	"common.dh.cn/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	host := beego.AppConfig.String("dataI_host")
	port := beego.AppConfig.String("dataI_port")
	name := beego.AppConfig.String("dataI_name")
	username := beego.AppConfig.String("dataI_username")
	password := beego.AppConfig.String("dataI_password")
	timezone := beego.AppConfig.DefaultString("dataI_timezone", "Asia/Shanghai")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s", username, password, host, port, name, url.QueryEscape(timezone))
	if host == "" {
		return
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("dataI", "mysql", connection)
	orm.SetMaxIdleConns("dataI", 30)
	orm.SetMaxOpenConns("dataI", 30)
	runmode := beego.AppConfig.DefaultString("runmode", "dev")
	if runmode == "dev" {
		orm.Debug = true
	}
	OrmLogger = utils.GetLogger("orm")
}

type DiBase struct {
}

func (m *DiBase) Orm() orm.Ormer {
	o := orm.NewOrm()
	err := o.Using("dataI")
	if err != nil {
		m.errReport(err)
		panic(err)
	}
	return o
}

func (m *DiBase) query(entity interface{}) orm.QuerySeter {
	return m.Orm().QueryTable(entity)
}

func (m *DiBase) create(entity interface{}) bool {
	mutable := reflect.ValueOf(entity).Elem()
	mutable.FieldByName("ObjectId").SetString(bson.NewObjectId().Hex())
	now := utils.NowTime()
	mutable.FieldByName("CreateTime").Set(reflect.ValueOf(now))
	mutable.FieldByName("UpdateTime").Set(reflect.ValueOf(now))
	_id, err := m.Orm().Insert(entity)
	if err != nil {
		m.errReport(err)
		return false
	} else {
		mutable.FieldByName("Id").SetInt(_id)
		return true
	}
}

func (m *DiBase) update(entity interface{}) bool {
	mutable := reflect.ValueOf(entity).Elem()
	now := utils.NowTime()
	mutable.FieldByName("UpdateTime").Set(reflect.ValueOf(now))
	_, err := m.Orm().Update(entity)
	if err != nil {
		m.errReport(err)
		return false
	} else {
		return true
	}
}

func (m *DiBase) delete(entity interface{}, args ...interface{}) bool {
	var err error
	total := len(args)
	if total == 1 {
		index := args[0]
		switch index.(type) {
		case string, *string:
			_, err = m.findByObjectID(entity, index.(string)).Delete()
		case int64, *int64:
			_, err = m.findByID(entity, index.(int64)).Delete()
		case map[string]interface{}:
			_, err = m.findByFilters(entity, index.(map[string]interface{})).Delete()
		default:
		}
		if err != nil {
			m.errReport(err)
			return false
		}
		return true
	} else if total == 2 {
		key := args[0]
		value := args[1]
		_key, ok := key.(string)
		if !ok {
			return false
		}
		_, err = m.findByFilter(entity, _key, value).Delete()
		if err != nil {
			m.errReport(err)
			return false
		}
		return true
	} else {
		return false
	}
}

func (m *DiBase) softDelete(entity interface{}, args ...interface{}) bool {
	var err error
	total := len(args)
	params := orm.Params{"status": -1}
	if total == 1 {
		index := args[0]
		switch index.(type) {
		case string, *string:
			_, err = m.findByObjectID(entity, index.(string)).Update(params)
		case int64, *int64:
			_, err = m.findByID(entity, index.(int64)).Update(params)
		case map[string]interface{}:
			_, err = m.findByFilters(entity, index.(map[string]interface{})).Update(params)
		default:
		}
		if err != nil {
			m.errReport(err)
			return false
		}
		return true
	} else if total == 2 {
		key := args[0]
		value := args[1]
		_key, ok := key.(string)
		if !ok {
			return false
		}
		_, err = m.findByFilter(entity, _key, value).Update(params)
		if err != nil {
			m.errReport(err)
			return false
		}
		return true
	} else {
		return false
	}
}

func (m *DiBase) find(entity interface{}, args ...interface{}) interface{} {
	total := len(args)
	var err error
	if total == 1 {
		index := args[0]
		switch index.(type) {
		case string, *string:
			err = m.findByObjectID(entity, index.(string)).One(entity)
		case int, *int:
			err = m.findByID(entity, index.(int64)).One(entity)
		case map[string]interface{}:
			err = m.findByFilters(entity, index.(map[string]interface{})).One(entity)
		case *orm.Condition:
			err = m.query(entity).SetCond(index.(*orm.Condition)).One(entity)
		default:
			return nil
		}
		if err != nil {
			m.errReport(err)
			return nil
		} else {
			return entity
		}
	} else if total == 2 {
		key := args[0]
		value := args[1]
		_key, ok := key.(string)
		if !ok {
			return nil
		}
		err = m.findByFilter(entity, _key, value).One(entity)
		if err != nil {
			m.errReport(err)
			return nil
		} else {
			return entity
		}
	} else {
		return nil
	}
}

func (m *DiBase) findByFilter(entity interface{}, key string, value interface{}) orm.QuerySeter {
	return m.query(entity).Filter(key, value)
}

func (m *DiBase) findByFilters(entity interface{}, filters interface{}) orm.QuerySeter {
	query := m.query(entity)
	switch value := filters.(type) {
	case map[string]interface{}:
		for k, v := range value {
			query = query.Filter(k, v)
		}
	case *orm.Condition:
		query = query.SetCond(value)
	default:
		return nil
	}

	return query
}

func (m *DiBase) findByID(entity interface{}, id int64) orm.QuerySeter {
	return m.query(entity).Filter("id", id)
}

func (m *DiBase) findByObjectID(entity interface{}, object_id string) orm.QuerySeter {
	return m.query(entity).Filter("object_id", object_id)
}

func (m *DiBase) count(entity interface{}, filters map[string]interface{}) int64 {
	result, err := m.findByFilters(entity, filters).Count()
	if err != nil {
		m.errReport(err)
		return 0
	}
	return result
}

func (m *DiBase) pager(entity interface{}, filters map[string]interface{}, page_size int64) (total int64, total_page int64) {
	total = m.count(entity, filters)
	total_page = int64(math.Ceil(float64(total) / float64(page_size)))
	return total, total_page
}

func (m *DiBase) pagerList(entity interface{}, page int64, page_size int64, filters map[string]interface{}) orm.QuerySeter {
	return m.findByFilters(entity, filters).Offset((page - 1) * page_size).Limit(page_size)
}

func (m *DiBase) errReport(err error) {
	if !strings.HasSuffix(err.Error(), "no row found") {
		OrmLogger.Error(err.Error())
	}
}
