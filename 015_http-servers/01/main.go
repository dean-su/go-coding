package main

import (
	"net/http"
	//"fmt"
	"html/template"
	"log"
)

type hp int

func (h hp) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	//fmt.Fprintln(w, "hello welcome")
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hp
	http.ListenAndServe(":8080", d)
}

