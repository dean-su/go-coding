package main

import (
	"text/template"
	"os"
	"log"
)
var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	xl := []string{"000", "111", "222"}

	err := tpl.Execute(os.Stdout, xl)
	if err != nil {
		log.Fatalln(err)
	}
}