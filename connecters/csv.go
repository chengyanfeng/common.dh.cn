package connecters

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"common.dh.cn/utils"
)

type CsvConnecter struct {
	file        string
	split       rune
	count       int
	column      int
	row         int
	head        []string
	column_type []string
	Data        []utils.P
}

func NewCsv(file string) *CsvConnecter {
	return &CsvConnecter{
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

func (a *CsvConnecter) SetSplit(split rune) {
	a.split = split
}

func (a *CsvConnecter) GetRow() int {
	return a.row
}

func (a *CsvConnecter) GetColumn() int {
	return a.column
}

func (a *CsvConnecter) GetHead() []string {
	return a.head
}

func (a *CsvConnecter) GetColumnType() []string {
	return a.column_type
}

func (a *CsvConnecter) getReader(file *os.File) *csv.Reader {
	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.TrailingComma = false
	reader.TrimLeadingSpace = true
	reader.Comma = a.split
	return reader
}

func (a *CsvConnecter) ReadHead() ([]string, error) {
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

func (a *CsvConnecter) ReadAll(use_head bool, parse_type bool) error {
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

func (a *CsvConnecter) Export() error {
	if a.file == "" {
		utils.Error("csv export error: no such file")
		return errors.New("csv export error: no such file")
	}
	if len(a.Data) == 0 {
		utils.Error("csv export error: data empty")
		return errors.New("csv export error: data empty")
	}
	file, err := os.Create(a.file)
	if err != nil {
		utils.Error("csv export error:" + err.Error())
		return errors.New("csv export error:" + err.Error())
	}
	defer file.Close()
	file.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(file)
	content := make([][]string, 0)
	head := make([]string, 0)
	for _, info := range a.Data {
		row := make([]string, 0)
		for key, value := range info {
			if len(content) == 0 {
				head = append(head, key)
			}
			row = append(row, utils.ToString(value))
		}
		if len(content) == 0 {
			content = append(content, head)
		}
		content = append(content, row)
	}
	writer.WriteAll(content)
	writer.Flush()
	return nil
}

func (a *CsvConnecter) parseType(record []string) []string {
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
