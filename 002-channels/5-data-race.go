package main 

func main() {
	var data int
	go func ()  {
		data ++
	}()

	println(data)
}