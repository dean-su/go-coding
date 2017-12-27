package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"time"
)

func chanFlow(left, right chan string, bufferLen int) {
	if bufferLen <= 0 {
		left <- <-right
	} else {
		for i := 0; i < bufferLen; i++ {
			left <- <-right
		}
	}
}

func genString() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	} else {
		return base64.URLEncoding.EncodeToString(b)
	}
}

func main() {

	nruntime := 100000
	chanBuffer := 1000
	result := make([]string, 0, 100)

	lastChan := make(chan string, chanBuffer)
	dataForChan := make([]string, 0, chanBuffer)

	for i := 0; i < chanBuffer; i++ {
		dataForChan = append(dataForChan, genString())
	}

	var left chan string = nil
	right := lastChan

	begin := time.Now()
	fmt.Println("begin at:", begin)
	for i := 0; i < nruntime; i++ {

		left, right = right, make(chan string, chanBuffer)
		go chanFlow(left, right, chanBuffer)
	}
	for i := 0; i < chanBuffer; i++ {
		right <- dataForChan[i]
	}

	for i := 0; i < chanBuffer; i++ {
		result = append(result, <-lastChan)
	}

	end := time.Now()
	fmt.Println("end   at:", end, time.Since(begin))
	//fmt.Println(result)
}
