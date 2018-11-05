package main

import (
	"fmt"
	"time"
)

func send1(ch chan int64) {
	time.Sleep(time.Duration(5) * time.Second)
	ch <- 5
}

func send2(ch chan int64) {
	time.Sleep(time.Duration(10) * time.Second)
	ch <- 10
}

func send3(ch chan int64) {
	time.Sleep(time.Duration(15) * time.Second)
	ch <- 15
}

func main() {
	ch := make(chan int64, 3)
	go send1(ch)
	go send2(ch)
	go send3(ch)
	var msg int64
	for i := 0; i < 3; i++ {
		select {
		case msg = <-ch:
		}
	}
	fmt.Println("msg is ", msg)
}
