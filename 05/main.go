package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id int
}

func produceJob(jobChan chan<- Job, period int) {
	defer close(jobChan)
	id := 1
	for start := time.Now(); time.Since(start) < time.Duration(period*1000)*time.Millisecond; id++ {
		jobChan <- Job{id: id}
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func worker(id int, jobChan <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobChan {
		fmt.Printf("Worker %d processing job id: %d\n", id, job.id)
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	}
}

func main() {
	jobChan := make(chan Job)
	go produceJob(jobChan, 10)

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, jobChan, &wg)
	}

	wg.Wait()
}
