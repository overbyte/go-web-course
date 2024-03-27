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

	// note a trailing slash will allow this route to handle any route beginning
	// with /alpha eg /alpha/ or /alpha/another/route
	// it seems like this is the better approach as it's possible to use this
	// and still add a handler for /alpha/beta (for example) 
	http.Handle("/alpha/", a)

	// whereas not using a trailing slash will make this route specific
	http.Handle("/beta", b)

	// this uses the default servemux by not passing a NewServeMux instance
	// into http.ListenAndServe. This uses a global DefaultServeMux with the 
	// handlers for the routes
	// NOTE not sure i like this as it's a global instance which isn't very
	// flexible. It would be very quick to implement for a small site tho
	http.ListenAndServe(":8080", nil)
}
