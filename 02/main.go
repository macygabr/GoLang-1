package main

import (
	"fmt"
	"sync"
	"time"
)

type cont struct {
	arr [5]int
}

func main() {
	var cont cont

	cont.arr = [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for i := 0; i < len(cont.arr); i++ {
		wg.Add(1)
		go cont.spinner(&wg, 1000*time.Millisecond, i)
	}

	wg.Wait()
	fmt.Print(cont.arr)
}

func (cont *cont) spinner(wg *sync.WaitGroup, delay time.Duration, index int) {
	defer wg.Done()
	fmt.Println("Start ", index)
	time.Sleep(delay)
	cont.arr[index] = cont.arr[index] * cont.arr[index]
}
