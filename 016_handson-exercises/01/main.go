package main

import (
	"net/http"
	"io"
)

func a(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "hello")
}

func b(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "welcome")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dean Su")
}

func main() {
	http.HandleFunc("/", a)
	http.HandleFunc("/dog/", b)
	http.HandleFunc("/me/", c)

	http.ListenAndServe(":8080", nil)
}

