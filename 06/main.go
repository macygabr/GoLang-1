package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	contextCancel()
	channel()
	timeout()
}

func contextCancel() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx, &wg)
	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\033[31mworker stopped\033[0m")
			return
		default:
			fmt.Println("\033[31mworker working\033[0m")
			time.Sleep(time.Second)
		}
	}
}

func channel() {
	var wg sync.WaitGroup
	stop := make(chan bool)
	wg.Add(1)
	go workerChannel(stop, &wg)
	time.Sleep(3 * time.Second)
	stop <- true
	wg.Wait()
}

func workerChannel(stop chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			time.Sleep(time.Second)
			fmt.Println("\033[32mworker stopped\033[0m")
			return
		default:
			fmt.Println("\033[32mworker working\033[0m")
			time.Sleep(time.Second)
		}
	}
}

func timeout() {
	var wg sync.WaitGroup
	result := make(chan string)

	wg.Add(1)
	go workerTimeout(result, &wg)

	for {
		msg := <-result
		fmt.Println("\033[33mworker :\033[0m", msg)
		if msg == "success" {
			break
		}
		time.Sleep(time.Second)
	}
	wg.Wait()
}

func workerTimeout(result chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for start := time.Now(); time.Since(start) < 3*time.Second; {
		result <- "work"
	}
	result <- "success"
}
