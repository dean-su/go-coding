package main

import (
	"fmt"
	"net/http"
)

type hp int

func (h hp) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Dean-Key", "this is from Dean Su")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
}

func main() {
	var hh hp
	http.ListenAndServe(":8080", hh)
}