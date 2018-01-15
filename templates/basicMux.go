package main

import (
	"net/http"
	"io"
)

type someroute int

func (sr someroute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello route1")
}

type someroute1 int

func (sr1 someroute1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello route2")
}

func main() {
		var sr1 someroute
		var t someroute1
		mux := http.NewServeMux()
		mux.Handle("/someroute", sr1)
		mux.Handle("/someroute1/", t) // Trailing forward slash will use the same handler t for any urls following /someroute1/
		http.ListenAndServe(":8080", mux)
}
