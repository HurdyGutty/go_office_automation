package worker

import (
	"fmt"

	"github.com/HurdyGutty/go_office_automation/pkg/readExcel"
	"github.com/HurdyGutty/go_office_automation/pkg/replace"
)

func worker(persons []readExcel.Person, temp_file, output_dir string, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		replace.Replace(persons[j], temp_file, fmt.Sprintf("%v/%d.docx", output_dir, j+1))
		results <- j
	}
}

func CreateWorker(input, temp_dir, output_dir string) {
	persons := readExcel.ReadExcel(input)
	var numJobs = len(persons)

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(persons, temp_dir, output_dir, jobs, results)
	}

	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 0; a < numJobs; a++ {
		<-results
	}
}
