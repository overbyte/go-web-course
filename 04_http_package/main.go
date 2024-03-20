package main

import (
	"html/template"
	"log"
	"net/http"
)

type mux int

func (m mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// req.Form returns a series of key,value pairs where key is a string and
	// value is a slice of string 
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
	// if we don't want to include the querystring in the form action, we can
	// use req.PostForm which only includes the POST, PUT or PATCH values
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var m mux
	http.ListenAndServe(":8080", m)
}
