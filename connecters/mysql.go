package connecters

import (
	"database/sql"
	"fmt"

	"common.dh.cn/utils"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConnecter struct {
	host     string
	port     string
	username string
	password string
}

func NewMysqlConnecter(host string, port string, username string, password string) *MysqlConnecter {
	return &MysqlConnecter{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

func (a *MysqlConnecter) buildConfig() utils.P {
	return utils.P{
		"host":     a.host,
		"port":     a.port,
		"username": a.username,
		"password": a.password,
		"fmt":      "mysl",
	}
}

func (a *MysqlConnecter) GetConnection() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)?charset=utf8", a.username, a.password, a.host, a.port)
}

func (a *MysqlConnecter) GetData(Sql string) *[]utils.P {
	db, err := sql.Open("mysql", a.GetConnection())
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT id,parent_id,qt_country_id,name FROM qt_areas limit 3")
	columns, err := rows.Columns()
	var result []utils.P
	for rows.Next() {
		column := make([](string), len(columns))
		row := utils.P{}
		args := make([]interface{}, 0)
		for k := range column {
			args = append(args, &column[k])
		}
		err = rows.Scan(args...)
		fmt.Println(column)
		for k, v := range columns {
			row[v] = column[k]
		}
		result = append(result, row)
	}
	return &result
}

func (a *MysqlConnecter) GetDataByJDBC(sql string) *[]utils.P {
	result, err := utils.JDBC(sql, a.buildConfig())
	if err != nil {
		utils.Error("jdbc mysql run error:" + err.Error())
		return nil
	}
	return utils.JsonDecodeArrays([]byte(result))
}
