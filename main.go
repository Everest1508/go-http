package main

import (
	"html/template"
	"log"
	"net/http"
)

type reqnres int

func (m reqnres) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(req.Method, req.Form)
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	var d reqnres
	http.ListenAndServe(":8080", d)
}
