package main

import (
	"fmt"

	"github.com/HurdyGutty/go_office_automation/pkg/readExcel"
	"github.com/HurdyGutty/go_office_automation/pkg/zipXML"
)

func main() {
	persons := readExcel.ReadExcel()
	fmt.Println(persons)
	zipXML.ZipFile("word_xml\\", "output.docx")
}
