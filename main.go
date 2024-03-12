package main

import (
	"fmt"

	"github.com/HurdyGutty/go_office_automation/pkg/readExcel"
	"github.com/HurdyGutty/go_office_automation/pkg/replace"
	// "github.com/HurdyGutty/go_office_automation/pkg/zipXML"
)

func worker(persons []readExcel.Person, temp_dir string, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		replace.Replace(persons[j], temp_dir, fmt.Sprintf("output/%d.docx", j+1))
		results <- j
	}
}

func main() {
	persons := readExcel.ReadExcel()
	length := len(persons)

	var numJobs = length
	temp_dir := "template/template_test.docx"

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(persons, temp_dir, jobs, results)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 0; a < numJobs; a++ {
		<-results
	}

}
