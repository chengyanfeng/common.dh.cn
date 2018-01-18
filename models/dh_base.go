package models

import (
	"fmt"
	"math"
	"net/url"
	"reflect"
	"time"

	"common.dh.cn/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

var OrmLogger *logrus.Logger

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	host := beego.AppConfig.String("mysql_host")
	port := beego.AppConfig.String("mysql_port")
	name := beego.AppConfig.String("mysql_name")
	username := beego.AppConfig.String("mysql_username")
	password := beego.AppConfig.String("mysql_password")
	timezone := beego.AppConfig.DefaultString("mysql_timezone", "Asia/Shanghai")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s", username, password, host, port, name, url.QueryEscape(timezone))
	if host == "" {
		return
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", connection)
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	runmode := beego.AppConfig.DefaultString("runmode", "dev")
	if runmode == "dev" {
		orm.Debug = true
	}
	OrmLogger = utils.GetLogger("orm")
}

type DhBase struct {
}

func (m *DhBase) Orm() orm.Ormer {
	o := orm.NewOrm()
	err := o.Using("default")
	if err != nil {
		panic(err)
	}
	return o
}

func (m *DhBase) query(entity interface{}) orm.QuerySeter {
	return m.Orm().QueryTable(entity)
}

func (m *DhBase) create(entity interface{}) bool {
	mutable := reflect.ValueOf(entity).Elem()
	mutable.FieldByName("ObjectId").SetString(bson.NewObjectId().Hex())
	now := time.Now()
	mutable.FieldByName("CreateTime").Set(reflect.ValueOf(now))
	mutable.FieldByName("UpdateTime").Set(reflect.ValueOf(now))
	_id, err := m.Orm().Insert(entity)
	if err != nil {
		return false
	} else {
		mutable.FieldByName("Id").SetInt(_id)
		return true
	}
}

func (m *DhBase) update(entity interface{}) bool {
	mutable := reflect.ValueOf(entity).Elem()
	now := time.Now()
	mutable.FieldByName("UpdateTime").Set(reflect.ValueOf(now))
	_, err := m.Orm().Update(entity)
	if err != nil {
		OrmLogger.Error(err)
		return false
	} else {
		return true
	}
}

func (m *DhBase) delete(entity interface{}, args ...interface{}) bool {
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
			OrmLogger.Error(err)
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
			OrmLogger.Error(err)
			return false
		}
		return true
	} else {
		return false
	}
}

func (m *DhBase) softDelete(entity interface{}, args ...interface{}) bool {
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
			OrmLogger.Error(err.Error())
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
			OrmLogger.Error(err.Error())
			return false
		}
		return true
	} else {
		return false
	}
}

func (m *DhBase) find(entity interface{}, args ...interface{}) interface{} {
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
			OrmLogger.Error(err.Error())
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
			OrmLogger.Error(err.Error())
			return nil
		} else {
			return entity
		}
	} else {
		return nil
	}
}

func (m *DhBase) findByFilter(entity interface{}, key string, value interface{}) orm.QuerySeter {
	return m.query(entity).Filter(key, value)
}

func (m *DhBase) findByFilters(entity interface{}, filters map[string]interface{}) orm.QuerySeter {
	query := m.query(entity)
	for k, v := range filters {
		query = query.Filter(k, v)
	}
	return query
}

func (m *DhBase) findByID(entity interface{}, id int64) orm.QuerySeter {
	return m.query(entity).Filter("id", id)
}

func (m *DhBase) findByObjectID(entity interface{}, object_id string) orm.QuerySeter {
	return m.query(entity).Filter("object_id", object_id)
}

func (m *DhBase) count(entity interface{}, filters map[string]interface{}) int64 {
	result, err := m.findByFilters(entity, filters).Count()
	if err != nil {
		OrmLogger.Error(err)
		return 0
	}
	return result
}

func (m *DhBase) pager(entity interface{}, filters map[string]interface{}, page_size int64) (total int64, total_page int64) {
	total = m.count(entity, filters)
	total_page = int64(math.Ceil(float64(total) / float64(page_size)))
	return total, total_page
}

func (m *DhBase) pagerList(entity interface{}, page int64, page_size int64, filters map[string]interface{}) orm.QuerySeter {
	return m.findByFilters(entity, filters).Offset((page - 1) * page_size).Limit(page_size)
}

func (m *DhBase) errReport(err interface{}) {
	OrmLogger.Error(err)
}
