package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	mymap := make(map[int]int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go add(&wg, &mu, &mymap, i)
	}

	wg.Wait()
	fmt.Println(mymap)
}

func add(wg *sync.WaitGroup, mu *sync.Mutex, mymap *map[int]int, i int) {
	defer wg.Done()
	defer mu.Unlock()

	mu.Lock()
	(*mymap)[i] = i
}
