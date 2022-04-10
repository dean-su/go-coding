package main

import (
	"sync"
	."fmt"
)

var (
	mutex sync.Mutex
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	Printf("Depositing amount: %d to account with balance:%d \n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	Printf("Withdrawing amount: %d from account with the balance: %d \n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}

func main() {
	println("Hello world mutex..")
	balance = 1000
	var wg  sync.WaitGroup
	wg.Add(2)
	go withdraw(700, &wg)
	go deposit(500, &wg)
	wg.Wait()

	println("New Balance: ", balance)
}