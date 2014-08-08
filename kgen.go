package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os/exec"
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
	url := r.FormValue("i")
	top := r.FormValue("top")
	bot := r.FormValue("bot")
	out, err := exec.Command("convert", url, "-font", "impact.ttf", "-stroke", "#000", "-fill", "#fff", "-pointsize", "33", "-gravity", "North", "-draw", "text 1,0 '"+top+"'", "-gravity", "South", "-draw", "text 1,0 '"+bot+"'", "-").CombinedOutput()
	if err != nil {
		fmt.Printf("%s", out)
	}
	encoder := base64.NewEncoder(base64.StdEncoding, w)
	encoder.Write(out)
	encoder.Close()
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/preview", previewHandler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))
	http.ListenAndServe(":8080", nil)
}

