package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", routeDefault)
	http.HandleFunc("/dog/", routeDog)
	http.HandleFunc("/me/", routeMe)

	http.ListenAndServe(":8080", nil)
}

func init() {
	tpl = template.Must(template.ParseFiles(
		"templates/index.gohtml",
		"templates/dog.gohtml",
		"templates/me.gohtml",
	))
}

func routeDefault(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func routeDog(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func routeMe(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "me.gohtml", "Allandt")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
