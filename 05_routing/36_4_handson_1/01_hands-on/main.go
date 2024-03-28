package main

import (
	"io"
	"net/http"
)

func routeDefault(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "default")
}

func routeDog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dog")
}

func routeMe(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Allandt")
}

func main() {
	http.HandleFunc("/", routeDefault)
	http.HandleFunc("/dog/", routeDog)
	http.HandleFunc("/me/", routeMe)

	http.ListenAndServe(":8080", nil)
}
