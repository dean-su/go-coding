package _1

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
	strt := struct {
		Name string
		Age int
		Boy bool
	}{
		"James Su",
		4,
		true,
	}

	err := tpl.Execute(os.Stdout, strt)
	if err != nil {
		log.Fatalln(err)
	}
}
