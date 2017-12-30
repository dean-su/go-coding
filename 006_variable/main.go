package main

import (
	"text/template"
	"os"
	"log"
)
var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `hello, I am from golang`)
	if err != nil{
		log.Fatalln(err)

	}
}