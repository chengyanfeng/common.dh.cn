package main

import (
	"io"
	"os"
	"bytes"
	"fmt"
	"flag"
	"strings"
	"text/template"
	"common.dh.cn/models"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TemplateData struct {
	TableName string
	ModelName string
	Fields []MateData
}

type MateData struct {
	Name string
	Type string
	Tag string
}

func main () {
	fmt.Println("welcome to use dh code ganerate:")
	//解析参数
    flag.Parse()
	var cmd string = flag.Arg(0);
	var table string = flag.Arg(1);
	if cmd == "ganerate_model" {
		code := ganerateModel(table)
		var file_name = fmt.Sprintf("models/%s.go",table)
		var f *os.File
		if checkFileIsExist(file_name) {
			_ = os.Remove(file_name)
		}
		f, _ = os.Create(file_name)
		io.WriteString(f, code)
		fmt.Printf("ganerate %s", file_name)
	} else if cmd == "ganerate_database" {
		ganerateDataBase()
	}
}

func ganerateDataBase() {
	sql := "show table status"
	var maps []orm.Params
	num, err := new(models.DhBase).Orm().Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		for _,info := range maps {
			table := info["Name"].(string)
			code := ganerateModel(table)
			var file_name = fmt.Sprintf("models/%s.go",table)
			var f *os.File
			if checkFileIsExist(file_name) {
				_ = os.Remove(file_name)
			}
			f, _ = os.Create(file_name)
			io.WriteString(f, code)
			fmt.Printf("ganerate %s\n", file_name)
		}
	}
}

func ganerateModel(table string) string {
	data := TemplateData{}
	data.TableName = table
	data.ModelName = strFirstToUpper(table)
	data.Fields = make([]MateData,0)
	sql := fmt.Sprintf("show full columns from %s", table)
	var maps []orm.Params
	num, err := new(models.DhBase).Orm().Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		for _,info := range maps {
			field_name := info["Field"].(string)
			field_type := info["Type"].(string)
			if field_name == "id" || field_name == "object_id" || field_name == "create_time" || field_name == "update_time" {
				continue
			}
			mate_data := MateData{}
			mate_data.Name = strFirstToUpper(field_name)
			mate_data.Tag = fmt.Sprintf("`json:\"%s\"`",field_name)
			if HasPrefix(field_type, "varchar") || HasPrefix(field_type, "char") || HasPrefix(field_type, "text") {
				mate_data.Type = "string"
			}
			if HasPrefix(field_type, "tinyint") {
				mate_data.Type = "int"
			}
			if HasPrefix(field_type, "timestamp") {
				mate_data.Type = "time.Time"
			}
			data.Fields = append(data.Fields, mate_data)
		}
	}
	tmpl, err := template.ParseFiles("generate/model.tpl")
	if err != nil {   
		panic(err)
	}

	var doc bytes.Buffer
	err = tmpl.Execute(&doc, data)
	if err != nil {
		panic(err)
	} 
	return doc.String()
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func strFirstToUpper(str string) string {
    temp := strings.Split(str, "_")
    var upperStr string
    for y := 0; y < len(temp); y++ {
        vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				vv[i] -= 32
				upperStr += string(vv[i]) // + string(vv[i+1])
			} else {
				upperStr += string(vv[i])
			}
		}
    }
    return upperStr
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}