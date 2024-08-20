package main

import (
	"fmt"
	"sync"
	"time"
)

type increment struct {
	val int
}

func main() {
	var value increment
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(3)
	go test(&wg, &mu, &value)
	go test(&wg, &mu, &value)
	go test(&wg, &mu, &value)
	wg.Wait()
	fmt.Println(value.val)
}

func test(wg *sync.WaitGroup, mu *sync.Mutex, val *increment) {
	defer wg.Done()
	defer mu.Unlock()
	mu.Lock()
	val.val = val.val + 1
	time.Sleep(time.Second)
}
