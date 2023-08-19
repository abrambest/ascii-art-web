package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func errorPage() {
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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	// mux.HandleFunc("/ascii-art", nil)
	log.Println("Server start: http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	log.Fatal(err)
}
