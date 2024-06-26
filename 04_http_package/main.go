package main

import (
	"html/template"
	"log"
	"net/http"
)

type mux int

func (m mux) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Context-Type", "text/html; charset=utf-8")
	res.Header().Set("Allandt-Custom", "All your base are belong to us")

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// req.Form returns a series of key,value pairs where key is a string and
	// value is a slice of string 
	tpl.ExecuteTemplate(res, "index.gohtml", req.Form)
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
