package main

import (
	"fmt"
	"text/template"
	"time"
	"os"
	"log"
)
var tpl *template.Template
var fm = template.FuncMap{
	"fDateTime" :	formatDatetime,
}

func formatDatetime(t time.Time) string {
	return t.Format(time.ANSIC)
}
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("done")
}

