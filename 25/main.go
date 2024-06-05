package main

import (
	"fmt"
	"time"
)

func MySleep(duration time.Duration) {
	<-time.After(2 * time.Second)
}

func main() {
	fmt.Println("Start")
	MySleep(2 * time.Second)
	fmt.Println("End after 2 seconds")
}
