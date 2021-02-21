package main

import (
	"fmt"
	"runtime"
)

func number_server(add <-chan int, sub <-chan int, read chan<- int) {
	var number = 0

	// This for-select pattern is one you will become familiar with...
	for {
		select {
		case nmb := <-add:
			number += nmb
		case nmb := <-sub:
			number -= nmb
		case read <- number:
			continue
		}
	}
}

func incrementer(add chan<- int, finished chan<- bool) {
	for j := 0; j < 1000001; j++ {
		add <- 1
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func decrementer(sub chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000+1; j++ {
		sub <- 1
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the remaining channels
	readCh := make(chan int)
	addCh := make(chan int)
	subCh := make(chan int)

	incrFinished := make(chan bool)
	decrFinished := make(chan bool)

	// TODO: Spawn the required goroutines
	go number_server(addCh, subCh, readCh)
	go incrementer(addCh, incrFinished)
	go decrementer(subCh, decrFinished)

	// TODO: block on finished from both "worker" goroutines
	<-incrFinished
	<-decrFinished

	fmt.Println("The magic number is:", <-readCh)
}
