package main

import (
	"ascii-web/asciiart"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func errorPage(w http.ResponseWriter, err error) {
	html, err := template.ParseFiles("assets/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = html.Execute(w, err)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func home(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("assets/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = html.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
func asciiHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("assets/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	font := r.FormValue("font")
	text := r.FormValue("text")
	res, err := asciiart.AsciiFunc(text, font)
	if err != nil {
		fmt.Println(err)
		errorPage(w, err)
		return
	}

	err = html.Execute(w, res)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	mux.HandleFunc("/ascii-art", asciiHandler)
	log.Println("Server start: http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	log.Fatal(err)
}
