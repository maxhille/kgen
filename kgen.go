package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// read dir, only take jpg
	fis, _ := ioutil.ReadDir("./img/")
	jpgs := make([]string, 0)
	for _, fi := range fis {
		if !strings.HasSuffix(fi.Name(), "jpg") {
			continue
		}
		jpgs = append(jpgs, fi.Name())
	}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, jpgs)
}

func previewHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("i")
	top := r.FormValue("top")
	bot := r.FormValue("bot")
	fmt.Fprintf(w, "%s / %s / %s", name, top, bot)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/preview", previewHandler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))
	http.ListenAndServe(":8080", nil)
}

