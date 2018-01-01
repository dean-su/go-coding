package main

import (
	"net/http"
	//"fmt"
	"html/template"
	"log"
	"net/url"
)

type hp int

func (h hp) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	//fmt.Fprintln(w, "hello welcome")
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method string
		Submissions url.Values
		URL    *url.URL
		Header http.Header
		Host   string
		ContentLength int64
	}{
		r.Method,
		r.Form,
		r.URL,
		r.Header,
		r.Host,
		r.ContentLength,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hp
	http.ListenAndServe(":8080", d)
}

