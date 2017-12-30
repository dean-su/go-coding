package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template
type family struct {
	Name string
	Age  int
	Boy  bool
}

func (f family) DAge() int  {
	return f.Age * 2
}
func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	/*strt := []family{
		family{"James Su", 4, true,},
		family{"Javon Su", 1, true},
		family{"Jinjin", 36, false},
	}*/
	strt := [] struct{
		Name string
		Age  int
		Boy  bool
	}{
		{"James Su", 4, true,},
		{"Javon Su", 1, true},
		{"Jinjin", 36, false},
	}
	strt1 := family{
		"Dean Su",
		36,
		false,
	}
	err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml", strt)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout,"index.gohtml", strt1)
	if err != nil {
		log.Fatalln(err)
	}
}
