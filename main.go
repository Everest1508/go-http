package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type reqnres int

func (m reqnres) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	data := struct {
		Method      string
		Submissions map[string][]string
		URL         *url.URL
		Header      http.Header
	}{
		r.Method,
		r.Form,
		r.URL,
		r.Header,
	}
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(r.Method, r.Form)
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	var d reqnres
	http.ListenAndServe(":8080", d)
}
