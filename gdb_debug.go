package main

import (
	"fmt"
	"time"
)

func counting(bus chan int) {
	defer close(bus)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		bus <- i
	}
}

func main() {
	bus := make(chan int)
	go counting(bus)
	for count := range bus {
		fmt.Println("count is %d", count)
	}
}
