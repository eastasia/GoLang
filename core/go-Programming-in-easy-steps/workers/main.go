package main

import (
	"fmt"
	"time"
)

func main() {

	numJobs := 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	go worker(1, jobs, results)
	go worker(2, jobs, results)
	go worker(3, jobs, results)

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		<-results
	}
	close(results)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Print("Worker ", id, " ran Job ", job)
		fmt.Printf(" ...%v x %v = %v \n", job, job, job*job)
		time.Sleep(1 * time.Second)
		results <- job
	}
}
