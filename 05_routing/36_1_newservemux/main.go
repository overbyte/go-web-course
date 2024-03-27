package main

import (
	"io"
	"net/http"
)

type routeA int

func (m routeA) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "this is route a")
}

type routeB int

func (m routeB) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "this is route b")
}

func main() {
	var a routeA
	var b routeB

	mux := http.NewServeMux()
	// note a trailing slash will allow this route to handle any route beginning
	// with /alpha eg /alpha/ or /alpha/another/route
	// it seems like this is the better approach as it's possible to use this
	// and still add a handler for /alpha/beta (for example) 
	mux.Handle("/alpha/", a)

	// whereas not using a trailing slash will make this route specific
	mux.Handle("/beta", b)

	http.ListenAndServe(":8080", mux)
}
