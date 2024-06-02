package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	id int
}

func produceJob(jobChan chan<- Job) {
	defer close(jobChan)
	for id := 1; true; id++ {
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
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number_of_workers>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])

	if err != nil || numWorkers <= 0 {
		fmt.Println("Invalid number of workers:", os.Args[1])
		return
	}

	jobChan := make(chan Job)
	go produceJob(jobChan)

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobChan, &wg)
	}

	wg.Wait()
}
