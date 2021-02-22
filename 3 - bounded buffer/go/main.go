package main

import (
	"fmt"
	"time"
)

func producer(addData chan<- int) {

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
		// TODO: push real value to buffer
		addData <- i
	}

}

func consumer(returData chan<- int) {

	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		j := returData //TODO: get real value from buffer
		fmt.Printf("[consumer]: %d\n", j)
		time.Sleep(50 * time.Millisecond)
	}

}

func main() {

	// TODO: make a bounded buffer
	buffer := make(chan int, 5)

	go consumer(buffer)
	go producer(buffer)

	select {}
}
