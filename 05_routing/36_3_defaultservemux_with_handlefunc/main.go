package main

import (
	"io"
	"net/http"
)

// even quicker implementation to set up functions for the routes and use
// http.HandleFunc to use functions rather than handlers
// this is the quickest and most dorty
func routeA(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "this is route a")
}

func routeB(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "this is route b")
}

func main() {
	// note a trailing slash will allow this route to handle any route beginning
	// with /alpha eg /alpha/ or /alpha/another/route
	// it seems like this is the better approach as it's possible to use this
	// and still add a handler for /alpha/beta (for example) 
	http.HandleFunc("/alpha/", routeA)

	// whereas not using a trailing slash will make this route specific
	http.HandleFunc("/beta", routeB)

	// this uses the default servemux by not passing a NewServeMux instance
	// into http.ListenAndServe. This uses a global DefaultServeMux with the 
	// handlers for the routes
	// NOTE not sure i like this as it's a global instance which isn't very
	// flexible. It would be very quick to implement for a small site tho
	http.ListenAndServe(":8080", nil)
}
