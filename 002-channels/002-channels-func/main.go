package main

import "fmt"

func main() {
	c := make(chan int)
	defer close(c)
	go recChan(c)
	sendChan(c)
	fmt.Println("Done")
}

func recChan(c chan<- int)  {
	c <- 33
}

func sendChan(c <-chan int)  {
	fmt.Printf("%v\n", <-c)
}
