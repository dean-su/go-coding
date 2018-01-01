package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	exitSig := make(chan int)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
		exitSig <- 0
	}()
	code := <-exitSig
	os.Exit(code)
}
