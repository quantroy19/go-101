package main

import "fmt"

func main() {
	ch := make(chan int)
	// ch <- 1
	go handle(ch)
	
	fmt.Println("Received:", <-ch)
}

func handle(ch chan int) {
	// do something
	ch <- 1
}