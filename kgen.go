package main

import (
	"html/template"
	"net/http"
	"io/ioutil"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// read dir, only take jpg
	fis, _ := ioutil.ReadDir("./img/")
	jpgs := make([]string, 0)
	for _, fi := range fis {
		if !strings.HasSuffix(fi.Name(),"jpg") {
			continue
		}
		jpgs = append(jpgs, fi.Name())
	}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, jpgs)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))
	http.ListenAndServe(":8080", nil)
}

