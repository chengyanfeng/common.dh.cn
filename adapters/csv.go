package adapters

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"common.dh.cn/utils"
)

type CsvAdapter struct {
	file        string
	split       rune
	count       int
	column      int
	row         int
	head        []string
	column_type []string
	Data        []utils.P
}

func NewCsvAdapter(file string) *CsvAdapter {
	return &CsvAdapter{
		file:        file,
		split:       ',',
		count:       0,
		column:      0,
		row:         0,
		head:        make([]string, 0),
		column_type: make([]string, 0),
		Data:        make([]utils.P, 0),
	}
}

func (a *CsvAdapter) SetSplit(split rune) {
	a.split = split
}

func (a *CsvAdapter) GetRow() int {
	return a.row
}

func (a *CsvAdapter) GetColumn() int {
	return a.column
}

func (a *CsvAdapter) GetHead() []string {
	return a.head
}

func (a *CsvAdapter) GetColumnType() []string {
	return a.column_type
}

func (a *CsvAdapter) getReader(file *os.File) *csv.Reader {
	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.TrailingComma = false
	reader.TrimLeadingSpace = true
	reader.Comma = a.split
	return reader
}

func (a *CsvAdapter) ReadHead() ([]string, error) {
	file, err := os.Open(a.file)
	if err != nil {
		utils.Error("csv open file error:" + err.Error())
		return nil, errors.New("csv open file error:" + err.Error())
	}
	defer file.Close()
	reader := a.getReader(file)
	record, err := reader.Read()
	if err == io.EOF {
		return nil, errors.New("csv file empty")
	}
	if err != nil {
		utils.Error("csv read error at row " + strconv.Itoa(a.row) + " :" + err.Error())
		return nil, errors.New("csv read error at row " + strconv.Itoa(a.row) + " :" + err.Error())
	}
	return record, nil
}

func (a *CsvAdapter) ReadAll(use_head bool, parse_type bool) error {
	file, err := os.Open(a.file)
	if err != nil {
		utils.Error("csv open file error:" + err.Error())
		return errors.New("csv open file error:" + err.Error())
	}
	defer file.Close()
	reader := a.getReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			utils.Error("csv read error at row " + strconv.Itoa(a.row) + " :" + err.Error())
			return errors.New("csv read error at row " + strconv.Itoa(a.row) + " :" + err.Error())
		}
		//第一行做为表头
		if len(a.head) == 0 {
			a.column = len(record)
			// 是否使用表头
			if use_head {
				a.head = record
			} else {
				for i := 0; i < a.column; i++ {
					a.head = append(a.head, "c"+strconv.Itoa(i))
				}
			}
		} else {
			//解析数据格式
			if parse_type && len(a.column_type) == 0 {
				a.column_type = a.parseType(record)
			}
			info := utils.P{}
			for i := 0; i < a.column; i++ {
				if a.column_type[i] == "date" {
					info[a.head[i]], _ = utils.ToDate(record[i])
				} else {
					info[a.head[i]] = record[i]
				}
			}
			a.row++
			a.Data = append(a.Data, info)
		}
	}
	return nil
}

func (a *CsvAdapter) parseType(record []string) []string {
	result := make([]string, 0)
	for _, data := range record {
		if utils.IsDate(data) {
			result = append(result, "date")
		} else if utils.IsInt(data) || utils.IsFloat(data) {
			result = append(result, "number")
		} else {
			result = append(result, "string")
		}
	}
	return result
}
