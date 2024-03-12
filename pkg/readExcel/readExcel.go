package readExcel

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

type Person map[string]interface{}

func openExcel() *excelize.File {
	f, err := excelize.OpenFile("../file_test/CV NHẤT.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return f
}

func closeExcel(f *excelize.File) {
	if err := f.Close(); err != nil {
		fmt.Println(err)
	}
}

func copyPerson(p Person) Person {
	new_person := Person{}
	for k, v := range p {
		new_person[k] = v
	}
	return new_person
}

func ReadExcel() []Person {
	f := openExcel()
	defer closeExcel(f)

	first_sheet := f.GetSheetName(0)

	person := Person{}

	rows, err := f.GetRows(first_sheet, excelize.Options{RawCellValue: true})
	if err != nil {
		fmt.Println(err)
	}
	keys := []string{}
	suffix_time := "(time)"
	for _, marker := range rows[0] {
		if strings.Contains(marker, suffix_time) {
			person[strings.TrimSuffix(marker, suffix_time)] = time.Now()
			keys = append(keys, strings.TrimSuffix(marker, suffix_time))
		} else {
			person[marker] = "A"
			keys = append(keys, marker)
		}
	}

	persons := []Person{}
	for _, row := range rows[1:] {
		new_person := copyPerson(person)
		for i, colCell := range row {
			if new_person[keys[i]] == time.Now() {
				new_person[keys[i]] = dateConv(colCell)
				continue
			}
			new_person[keys[i]] = colCell
		}
		persons = append(persons, new_person)
	}
	return persons
}

func dateConv(colCell string) time.Time {
	float_date, err := strconv.ParseFloat(colCell, 64)
	if err != nil {
		fmt.Println(err)
	}
	time, err := excelize.ExcelDateToTime(float_date, false)
	if err != nil {
		fmt.Println(err)
	}
	return time
}
