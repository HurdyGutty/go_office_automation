package main

import (
	"fmt"

	"github.com/HurdyGutty/go_office_automation/pkg/readExcel"
)

func main() {
	persons := readExcel.ReadExcel()
	fmt.Println(persons)
}
