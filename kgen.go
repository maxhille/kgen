//  kgen - meme generator
//
//  Copyright 2014, Max Hille <mh@lambdasoup.com>
//
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
	filename := makeFilename(out)
	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
	}
	file.Write(out)
	file.Close()
	http.Redirect(w, r, filename, http.StatusSeeOther)
}

func makeFilename(bs []byte) string {
	hash8 := fmt.Sprintf("%x", sha1.Sum(bs))[0:8]
	return fmt.Sprintf("pub/%s.jpg", hash8)
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
	http.ListenAndServe("localhost:1887", nil)
}

