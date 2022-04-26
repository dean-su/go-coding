package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("hello world")
	server := socketio.NewServer(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	server.On("connection", func(so socketio.Socket)  {
		log.Println("New Connection")
	})

	http.Handle("/socket.io", server)
	log.Fatal(http.ListenAndServe(":5000", nil))
}