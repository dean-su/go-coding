package main

import (
	"text/template"
	"os"
	"log"
)
var tpl *template.Template
type person struct {
	First string
	Last string
	Age int
}
func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {
	s := []string {"hello", "golang", "channel", "goroutine"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", s)
	if err != nil{
		log.Fatalln(err)

	}
	m := map[string]string {
		"one" : "hello",
		"two" : "golang",
		"three" : "channel",
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", m)
	if err != nil{
		log.Fatalln(err)

	}

	st1 := person{"James", "Su", 4}
	st2 := person{"Jovon", "Su", 1}
	st := []person{st1, st2}
	t, err := template.New("todoc").Parse("My name is {{range .}} {{.First}} {{.Last}} and age is {{.Age}} {{end}}")
	if err != nil {
		log.Fatalln(err)
	}

	err = t.Execute(os.Stdout, st)
	if err != nil{
		log.Fatalln(err)

	}

}