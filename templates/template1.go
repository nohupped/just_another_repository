package main

import (
	"text/template"
	"os"
)

func main() {

	var tpl *template.Template
	tpl = template.Must(template.ParseFiles("template1.gohtml"))
	//sages := []string{
	//	"gandhi",
	//	"mlk",
	//	"buddha",
	//	"jesus",
	//	"muhammad",
	//}
	sages1 := map[string]string{
		"india": "gandhi",
		"buddhism": "buddha",
	}
	html, err := os.Create("template1.html")
	if err != nil {
		panic(err)
	}
	//err = tpl.Execute(html, sages)
	err = tpl.Execute(html, sages1)
	if err != nil {
		panic(err)
	}

}
