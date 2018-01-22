package connecters

import (
	"errors"
	"strconv"

	"common.dh.cn/utils"
	"github.com/tealeg/xlsx"
)

type ExcelConnecter struct {
	file        string
	sheet       []string
	head        map[string][]string
	column_type map[string][]string
	Data        map[string][]utils.P
}

func NewExcel(file string) *ExcelConnecter {
	return &ExcelConnecter{
		file:        file,
		sheet:       make([]string, 0),
		head:        make(map[string][]string, 0),
		column_type: make(map[string][]string, 0),
		Data:        make(map[string][]utils.P, 0),
	}
}

func (a *ExcelConnecter) GetSheet() []string {
	return a.sheet
}

func (a *ExcelConnecter) GetHead(sheet string) []string {
	return a.head[sheet]
}

func (a *ExcelConnecter) GetColumnType(sheet string) []string {
	return a.column_type[sheet]
}

func (a *ExcelConnecter) ReadAll(use_head bool, parse_type bool) error {
	file, err := xlsx.OpenFile(a.file)
	if err != nil {
		utils.Error("xlsx read error:" + err.Error())
		return errors.New("xlsx read error:" + err.Error())
	}
	for key, sheet := range file.Sheet {
		a.sheet = append(a.sheet, key)
		a.head[key] = make([]string, 0)
		a.column_type[key] = make([]string, 0)
		a.Data[key] = make([]utils.P, 0)
		for _, row := range sheet.Rows {
			//第一行做为表头
			if len(a.head[key]) == 0 {
				// 是否使用表头
				for i := 0; i < len(row.Cells); i++ {
					if use_head {
						a.head[key] = append(a.head[key], row.Cells[i].String())
					} else {
						a.head[key] = append(a.head[key], "c"+strconv.Itoa(i))
					}
				}
			} else {
				info := utils.P{}
				for i := 0; i < len(row.Cells); i++ {
					//解析数据格式
					if parse_type && len(a.column_type[key]) == 0 {
						a.column_type[key] = append(a.column_type[key], a.parseType(row.Cells[i].String()))
					}
					info[a.head[key][i]] = row.Cells[i].String()
				}
				a.Data[key] = append(a.Data[key], info)
			}
		}
	}
	return nil
}

func (a *ExcelConnecter) parseType(data string) string {
	if utils.IsDate(data) {
		return "date"
	} else if utils.IsInt(data) || utils.IsFloat(data) {
		return "number"
	} else {
		return "string"
	}
}
