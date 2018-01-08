package main

import (
	"net/http"
	"io"
	"html/template"
	"log"
)

func a(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "hello")
}

func b(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "welcome")
}

func c(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "Dean Su")
	var tpl *template.Template
	tpl = template.Must(template.ParseFiles("welcome.gohtml"))
	/*tpl, err := template.ParseFiles("welcome.gohtml")
	if err != nil {
		log.Fatalln(err)
	}*/

	err := tpl.ExecuteTemplate(w, "welcome.gohtml", "Dean Su")
	if err != nil {
		log.Fatalln(err)
	}


}

func main() {
/*	http.HandleFunc("/", a)
	http.HandleFunc("/dog/", b)
	http.HandleFunc("/me/", c)*/

	http.Handle("/", http.HandlerFunc(a))
	http.Handle("/dog/", http.HandlerFunc(b))
	http.Handle("/me/", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}

