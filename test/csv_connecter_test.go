package test

import (
	"fmt"
	"testing"

	"common.dh.cn/connecters"
	"common.dh.cn/utils"
)

func TestCsvRead(t *testing.T) {
	fmt.Println("begin csv connecter testing......")
	csv := connecters.NewCsv("../samples/sample.csv")
	err := csv.ReadAll(true, true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(csv.GetHead())
	fmt.Println(csv.GetColumnType())
	fmt.Println(csv.GetColumn())
	fmt.Println(csv.GetRow())
	fmt.Println(csv.Data[0])
}

func TestCsvExport(t *testing.T) {
	fmt.Println("begin export csv testing......")
	csv := connecters.NewCsv("../samples/export.csv")
	csv.Data = []utils.P{
		utils.P{
			"id":   1,
			"name": "test",
			"time": "2017-1-1",
		},
		utils.P{
			"id":   2,
			"name": "test2",
			"time": "2017-1-2",
		},
	}
	err := csv.Export()
	if err != nil {
		t.Fatal(err)
	}
}
