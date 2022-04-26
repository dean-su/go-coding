package main

import (
	"encoding/json"
	"fmt"
	"go-coding/026_protocol_buffers/model"
	"io/ioutil"
	_ "io/ioutil"
	"log"

	_ "google.golang.org/protobuf/proto"
	//"github.com/golang/protobuf/proto"
)

func main() {

	fmt.Println("hello world")
	book := &model.Book{
		Id:    1,
		Title: "Master protobuf",
		Authors: []*model.Author{
			{Id: 1, Name: "Javon Su"},
		},
		Category: model.Category_Novel,
	}
	
	data, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	ioutil.WriteFile("book.json", data, 0600)

	//data, err = proto.Marshall(book)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(data)

	//ioutil.WriteFile("book.protobuf", data, 0600)

	//decode the data from


	// d := &Person {
	// 	Name: "Javon",
	// 	Age: 6,
	// }

	// data, err := proto.Marshall(d)
	// if err != nil {
	// 	log.Fatal("Marshalling err:", err)
	// }
	// fmt.Println(data)

}
