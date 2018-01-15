package main

import (
	"net/http"
	"io"
	"os"
)

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<image src="/resources/dog.jpg"> `)
}

func dogpic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("assets/dog.jpg")
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
	}
	http.ServeContent(w, r, f.Name(), fi.ModTime(),f)
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", dog)
	http.HandleFunc("/dog", dogpic)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}


