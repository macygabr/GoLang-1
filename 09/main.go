package main

import (
	"fmt"
	"sync"
)

func main() {
	data := [6]int{2, 3, 5, 7, 11, 13}

	var wg sync.WaitGroup
	arr := make(chan int, 5)
	pack := make(chan int, 5)
	wg.Add(3)
	go generator(&wg, arr, data)
	go squares(&wg, arr, pack)
	go listener(&wg, pack)
	wg.Wait()
}

func generator(wg *sync.WaitGroup, arr chan int, data [6]int) {
	defer wg.Done()
	for i := 0; i < len(data); i++ {
		arr <- data[i]
	}
	close(arr)
}

func squares(wg *sync.WaitGroup, arr chan int, pack chan int) {
	defer wg.Done()
	for x := range arr {
		pack <- 2 * x
	}
	close(pack)
}

func listener(wg *sync.WaitGroup, pack chan int) {
	defer wg.Done()
	for x := range pack {
		fmt.Println(x)
	}
}
