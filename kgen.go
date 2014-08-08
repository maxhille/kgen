package main

import (
	"encoding/base64"
	"html/template"
	"io/ioutil"
	"net/http"
	"os/exec"
	"crypto/sha1"
	"os"
	"fmt"
	"log"
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
	url := r.FormValue("url")
	top := r.FormValue("top")
	bot := r.FormValue("bot")
	out := render(url, top, bot)
	encoder := base64.NewEncoder(base64.StdEncoding, w)
	encoder.Write(out)
	encoder.Close()
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	top := r.FormValue("top")
	bot := r.FormValue("bot")
	out := render(url, top, bot)
	filename := fmt.Sprintf("pub/%x.jpg", sha1.Sum(out))
	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
	}
	file.Write(out)
	file.Close()
	http.Redirect(w, r, filename, http.StatusSeeOther)
}

func render(url, top, bot string) (out []byte) {
	out, _ = exec.Command("convert", url, "-font", "impact.ttf", "-stroke", "#000", "-fill", "#fff", "-pointsize", "33", "-gravity", "North", "-draw", "text 1,0 '"+top+"'", "-gravity", "South", "-draw", "text 1,0 '"+bot+"'", "-").CombinedOutput()
	return
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/preview", previewHandler)
	http.HandleFunc("/create", createHandler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))
	http.ListenAndServe(":8080", nil)
}

