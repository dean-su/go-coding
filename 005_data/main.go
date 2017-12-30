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
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 20)
	if err != nil{
		log.Fatalln(err)

	}
}