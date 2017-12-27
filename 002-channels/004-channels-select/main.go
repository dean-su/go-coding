package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan bool)

	go sendChan(ch1, ch2, ch3)

	recChan(ch1, ch2, ch3)

	fmt.Println("Done")
}

func recChan(c1, c2 chan int, c3 chan bool)  {
	for {

		select {
		case msg1 := <-c1:
			fmt.Println("receive msg:", msg1)
		case msg2 := <-c2:
			fmt.Println("receive msg:", msg2)
		case sig := <-c3:
			fmt.Println("receive quit signal:", sig)
			return
		default:
			fmt.Println("no receive signal")
		}
	}
}

func sendChan(c1, c2 chan int, c3 chan bool)  {
	for i := 0; i < 10; i++ {
		if i%2 == 0{
			c1 <- i
		}else {
			c2 <- i
		}
	}

	close(c3)
}