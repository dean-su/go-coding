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
	chanBuffer := 1000
	result := make([]int, 0, 100)

	lastChan := make(chan int, chanBuffer)

	var left chan int = nil
	right := lastChan

	begin := time.Now()
	fmt.Println("begin at:", begin)
	for i := 0; i < nruntime; i++ {
		left, right = right, make(chan int, chanBuffer)
		go chanFlow(left, right, chanBuffer)
	}
	fmt.Println("Num Chan:", runtime.NumGoroutine())
	for i := 0; i < chanBuffer; i++ {
		right <- 0
	}

	for i := 0; i < chanBuffer; i++ {
		result = append(result, <-lastChan)
	}

	end := time.Now()
	fmt.Println("end   at:", end, time.Since(begin))
	//fmt.Println(result)
	fmt.Println("Num Chan:", runtime.NumGoroutine())
}
