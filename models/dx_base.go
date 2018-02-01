package models

import (
	"fmt"
	"math"
	"net/url"
	"reflect"
	"strings"
	"time"

	"common.dh.cn/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	host := beego.AppConfig.String("dataX_host")
	port := beego.AppConfig.String("dataX_port")
	name := beego.AppConfig.String("dataX_name")
	username := beego.AppConfig.String("dataX_username")
	password := beego.AppConfig.String("dataX_password")
	timezone := beego.AppConfig.DefaultString("dataX_timezone", "Asia/Shanghai")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s", username, password, host, port, name, url.QueryEscape(timezone))
	if host == "" {
		return
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("dataX", "mysql", connection)
	orm.SetMaxIdleConns("dataX", 30)
	orm.SetMaxOpenConns("dataX", 30)
	runmode := beego.AppConfig.DefaultString("runmode", "dev")
	if runmode == "dev" {
		orm.Debug = true
	}
	OrmLogger = utils.GetLogger("orm")
}

type DxBase struct {
}

func (m *DxBase) Orm() orm.Ormer {
	o := orm.NewOrm()
	err := o.Using("dataX")
	if err != nil {
		m.errReport(err)
		panic(err)
	}
	return o
}

func (m *DxBase) query(entity interface{}) orm.QuerySeter {
	return m.Orm().QueryTable(entity)
}

func (m *DxBase) create(entity interface{}) bool {
	mutable := reflect.ValueOf(entity).Elem()
	mutable.FieldByName("ObjectId").SetString(bson.NewObjectId().Hex())
	now := time.Now()
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

func (m *DxBase) update(entity interface{}) bool {
	mutable := reflect.ValueOf(entity).Elem()
	now := time.Now()
	mutable.FieldByName("UpdateTime").Set(reflect.ValueOf(now))
	_, err := m.Orm().Update(entity)
	if err != nil {
		m.errReport(err)
		return false
	} else {
		return true
	}
}

func (m *DxBase) delete(entity interface{}, args ...interface{}) bool {
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

func (m *DxBase) softDelete(entity interface{}, args ...interface{}) bool {
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

func (m *DxBase) find(entity interface{}, args ...interface{}) interface{} {
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

func (m *DxBase) findByFilter(entity interface{}, key string, value interface{}) orm.QuerySeter {
	return m.query(entity).Filter(key, value)
}

func (m *DxBase) findByFilters(entity interface{}, filters map[string]interface{}) orm.QuerySeter {
	query := m.query(entity)
	for k, v := range filters {
		query = query.Filter(k, v)
	}
	return query
}

func (m *DxBase) findByID(entity interface{}, id int64) orm.QuerySeter {
	return m.query(entity).Filter("id", id)
}

func (m *DxBase) findByObjectID(entity interface{}, object_id string) orm.QuerySeter {
	return m.query(entity).Filter("object_id", object_id)
}

func (m *DxBase) count(entity interface{}, filters map[string]interface{}) int64 {
	result, err := m.findByFilters(entity, filters).Count()
	if err != nil {
		m.errReport(err)
		return 0
	}
	return result
}

func (m *DxBase) pager(entity interface{}, filters map[string]interface{}, page_size int64) (total int64, total_page int64) {
	total = m.count(entity, filters)
	total_page = int64(math.Ceil(float64(total) / float64(page_size)))
	return total, total_page
}

func (m *DxBase) pagerList(entity interface{}, page int64, page_size int64, filters map[string]interface{}) orm.QuerySeter {
	return m.findByFilters(entity, filters).Offset((page - 1) * page_size).Limit(page_size)
}

func (m *DxBase) errReport(err error) {
	if !strings.HasSuffix(err.Error(), "no row found") {
		OrmLogger.Error(err.Error())
	}
}
