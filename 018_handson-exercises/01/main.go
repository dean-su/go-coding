package main

import (
	"net/http"
	"io"
	"html/template"
)

func foo(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request)  {
	var tpl *template.Template
	tpl = template.Must(template.ParseFiles("dog.gohtml"))

	err := tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		http.Error(w, "ExecuteTemplate error", http.StatusSeeOther)
	}
}

func showPic(w http.ResponseWriter, r *http.Request)  {
	http.ServeFile(w, r, "merry.jpg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/merry.jpg", showPic)
	http.ListenAndServe(":8080", nil)
}