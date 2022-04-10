package main
import (
	."fmt"
	"time"
)

func compute(val int) {
	for i := 0; i < val; i++ {
		time.Sleep(time.Second)
		Println(i)
	}
}

func main() {
	Println("Concurrency...")
	// go compute(5)
	// go compute(5)
	go func() {
		println("executing the anonymous func...")
	}()
	var in string
	Scanln(&in)
}