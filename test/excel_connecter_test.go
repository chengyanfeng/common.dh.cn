package test

import (
	"fmt"
	"testing"

	"common.dh.cn/connecters"
)

func TestExcelRead(t *testing.T) {
	fmt.Println("begin csv connecter testing......")
	xlsx := connecters.NewExcel("../samples/sample.xlsx")
	err := xlsx.ReadAll(true, true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(xlsx.GetSheet())
}
