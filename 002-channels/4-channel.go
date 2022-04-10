package main

import (
	"math/rand"
)

func calculateValue(vals chan int) {
	val := rand.Intn(10)
	println("random value: {}", val)
	vals <- val 
}

func main() {
	println("Go channel")
	vals := make(chan int)
	defer close(vals)

	go calculateValue(vals)

	val := <- vals
	println(val)
}