package main

import (
	"fmt"
	"text/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		log.Fatalln(err)
	}
	f, err1:= os.Create("index.html")
	defer f.Close()
	if err1 != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(f, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Done")
}