package models

import (
	"time"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRPostgres)
	host := beego.AppConfig.String("postgres_host")
	port := beego.AppConfig.String("postgres_port")
	name := beego.AppConfig.String("postgres_name")
	username := beego.AppConfig.String("postgres_username")
	password := beego.AppConfig.String("postgres_password")
	connection := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", username, password, name, host, port)
	orm.RegisterDataBase("default", "postgres", connection)
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterModel(new(User))
	orm.Debug = true
}