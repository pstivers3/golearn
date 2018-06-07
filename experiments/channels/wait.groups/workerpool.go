package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id		 int
	randomno int
}

type Result struct {
	job			Job
	sumofdigits int
}

var jobs = make (chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit = no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Results {job, digits(job.randomno)}
		results <- output
	}
	wg.Done
}

func createWorkerPool(noOfWorkers in) {
	var wg sync.WaitGroup
	for i := 0 i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
	}
	close(jobs)
}


