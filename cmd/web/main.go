package main

import (
	"ascii-web/asciiart"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Error struct {
	Error string
	Code  int
}

func errorPage(w http.ResponseWriter, error string, code int) {
	htmlFiles := []string{
		"./assets/error.html",
		"./assets/base.layout.html",
		"./assets/footer.html",
	}

	html, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		fmt.Println("1")
		return
	}
	w.WriteHeader(code)
	err = html.ExecuteTemplate(w, "error.html", Error{
		Error: error,
		Code:  code,
	})

	if err != nil {
		fmt.Println("2")
		http.Error(w, "Internal Server Error", 500)
		return
	}

}

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		errorPage(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	htmlFiles := []string{
		"./assets/home.html",
		"./assets/footer.html",
		"./assets/base.layout.html",
	}
	if r.Method != http.MethodGet {
		errorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	html, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		errorPage(w, "Internal Server Error", 500)
		return
	}
	err = html.Execute(w, nil)
	if err != nil {
		errorPage(w, "Internal Server Error", 500)
		return
	}

}
func asciiHandler(w http.ResponseWriter, r *http.Request) {
	htmlFiles := []string{
		"./assets/home.html",
		"./assets/footer.html",
		"./assets/base.layout.html",
	}
	if r.Method != http.MethodPost {
		errorPage(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	html, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		errorPage(w, "Internal Server Error", 500)
		return
	}
	font := r.FormValue("font")
	text := r.FormValue("text")
	res, err := asciiart.AsciiFunc(text, font)
	if err != nil {
		err = html.Execute(w, err)
		if err != nil {
			errorPage(w, "Internal Server Error", 500)
			return
		}
	}

	err = html.Execute(w, res)
	if err != nil {
		errorPage(w, "Internal Server Error", 500)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	mux.HandleFunc("/ascii-art", asciiHandler)
	log.Println("Server start: http://localhost:3030")
	err := http.ListenAndServe("localhost:3030", mux)
	log.Fatal(err)
}
