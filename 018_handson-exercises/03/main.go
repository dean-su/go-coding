package main

import (
	"net/http"
	"log"
	"html/template"
)
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	//fs := http.FileServer(http.Dir("public"))
	//http.Handle("/pics/", fs)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))

	http.HandleFunc("/", dogs)
	http.ListenAndServe(":8080", nil)
}

func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
