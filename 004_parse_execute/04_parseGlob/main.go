package main

import (
	"fmt"
	"text/template"
	"log"
	"os"
)
var tpl *template.Template
func init()  {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}
func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("done")

}