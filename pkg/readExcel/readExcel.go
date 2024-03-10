package readExcel

import (
	"fmt"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

type Person struct {
	Name        string
	DateOfBirth time.Time
	Address     string
	Phone       string
}

func ReadExcel() []Person {
	f, err := excelize.OpenFile("../file_test/CV NHẤT.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	first_sheet := f.GetSheetName(0)

	rows, err := f.GetRows(first_sheet, excelize.Options{RawCellValue: true})
	if err != nil {
		fmt.Println(err)
	}
	persons := []Person{}
	for _, row := range rows[1:] {
		person := Person{}
		for i, colCell := range row {
			switch i {
			case 1:
				float_date, err := strconv.ParseFloat(colCell, 64)
				if err != nil {
					fmt.Println(err)
				}
				time, err := excelize.ExcelDateToTime(float_date, false)
				if err != nil {
					fmt.Println(err)
				}
				person.DateOfBirth = time
			case 0:
				person.Name = colCell
			case 2:
				person.Address = colCell
			case 3:
				person.Phone = colCell
			default:
				fmt.Println("Error")
			}
		}
		persons = append(persons, person)
	}
	return persons
}
