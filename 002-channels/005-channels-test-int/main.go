package main

import (
	"fmt"
	"time"
	"runtime"
)

func chanFlow(left, right chan int, bufferLen int) {
	if bufferLen <= 0 {
		left <- 1 + <-right
	} else {
		for i := 0; i < bufferLen; i++ {
			left <- 1 + <-right
		}
	}
}

func main() {
	fmt.Println("Num Chan:", runtime.NumGoroutine())
	nruntime := 100000
	lastChan := make(chan int)

	var left chan int = nil
	right := lastChan

	begin := time.Now()
	fmt.Println("begin at:", begin)
	for i := 0; i < nruntime; i++ {
		left, right = right, make(chan int)
		go chanFlow(left, right, 0)
	}

	right <- 0
	result := <-lastChan

	end := time.Now()
	fmt.Println("end   at:", end, time.Since(begin))
	fmt.Println(result)
	fmt.Println("Num Chan:", runtime.NumGoroutine())
}