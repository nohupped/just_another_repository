package main

import (
	"net/http"
	"html/template"
)

type somehandler int

func (s somehandler) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
		err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	http_form_and_template_tpl.ExecuteTemplate(w, "http_form_and_template.gohtml", req.Form)

}

var http_form_and_template_tpl *template.Template

func main() {

	http_form_and_template_tpl = template.Must(template.ParseFiles("http_form_and_template.gohtml"))
	var s somehandler
	http.ListenAndServe(":8080", s)

	}
