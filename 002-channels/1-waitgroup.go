package main

import (
	. "fmt"
	"sync"
	"time"
)
func myfunc(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	Println("finished goroutine")
	wg.Done()
}

func main() {
	println("Go waitgroup...")
	var wg sync.WaitGroup
	wg.Add(1)
	//go myfunc(&wg)
	go func() {
		println("Inside the goroutine")
		wg.Done()
	}()
	wg.Wait()
	println("Well done...")

}