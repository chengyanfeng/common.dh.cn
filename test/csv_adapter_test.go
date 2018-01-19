package test

import (
	"fmt"
	"testing"

	"common.dh.cn/adapters"
)

func TestCsvAdapter(t *testing.T) {
	fmt.Println("begin csv adapter testing......")
	csv := adapters.NewCsvAdapter("../samples/sample.csv")
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
