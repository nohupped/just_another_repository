package main

import (
	"net/http"
	"fmt"
)

type shit int

func (s shit)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "This satisfies interface 'Handler' from package net/http because it has 'ServeHTTP(w http.ResponseWriter, r *http.Request)' method attached to the underlying type.")

}

func main() {
	var s shit
	http.ListenAndServe(":8080", s)
}
