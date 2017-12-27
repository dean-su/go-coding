package main

import (
	"fmt"
	"sync"
	)

func main()  {
	c := make(chan int, 2)
	go func() {
		c <- 31
		c <- 32
		close(c)
	}()


	fmt.Println(<-c, <-c)

	cr := make(<- chan int, 1)
	cs := make(chan <- int, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		cs <- 33
		cr = c
		close(cs)

	}()
	fmt.Printf("cr:%v\t %T\n", cr, cr)
	fmt.Printf("cr:%v\t %T\n", cs , cs)
	wg.Wait()
}
