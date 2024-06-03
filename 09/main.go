package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	pack := make(chan int)

	wg.Add(1)
	go NurlanTheBest(&wg, &pack)

	wg.Add(1)
	go macygabr(&wg, &pack)

	wg.Wait()
}

func NurlanTheBest(wg *sync.WaitGroup, pack *chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		(*pack) <- i
	}
}

func macygabr(wg *sync.WaitGroup, pack *chan int) {
	defer wg.Done()
	// fmt.Println(<-pack)
}
