package main

import (
	"net/http"
	"io"
	//"os"
)

func merry(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="merry.jpg">`)
}

func merryPic(w http.ResponseWriter, r *http.Request)  {
	http.ServeFile(w, r, "merry.jpg")
/*	f, err := os.Open("merry.jpg")
	if err != nil {
		http.Error(w, "picture not found", 404)
		return
	}
	defer f.Close()
	//io.Copy(w, f)
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)*/

}

func main() {
	http.HandleFunc("/", merry)
	http.HandleFunc("/merry.jpg", merryPic)
	http.ListenAndServe(":8080", nil)
}
