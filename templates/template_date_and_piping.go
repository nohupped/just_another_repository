package main

import (
	"text/template"
	"time"
	"os"
	"math"
)

var template_date *template.Template

func monthDayYear(t time.Time) string {
	return t.Format(time.Kitchen)
}
var format_month_and_math_fns = template.FuncMap{
	"fdateMDY": monthDayYear,
	"fdouble": funcdouble,
	"fsqr": funcsqr,
		"fsqrt": funcsqrt,
}

func funcdouble(i int) int {
	return i * 2
}
func funcsqr(i int) float64 {
	return math.Pow(float64(i), float64(2))
}
func funcsqrt(i float64) float64 {
	return math.Sqrt(float64(i))
}


func main() {
	
	somefile, _ := os.Create("template_date_and_piping.html")

	// Make sure the struct members inside the struct that wraps up datetime and value
	// to pass to math functions are exported or it won't be visible inside the template
	// to be parsed.
	//
	data := struct {
		Datetime time.Time
		Valtodomath int
	}{time.Now(), 10}
	
	template_date = template.Must(template.New("").Funcs(format_month_and_math_fns).ParseFiles("template_date_and_piping.gohtml"))
	template_date.ExecuteTemplate(somefile, "template_date_and_piping.gohtml", data)

}
