package main

import "net/http"

import "html/template"

var templ = template.Must(template.ParseGlob("delivery/web/templates/*"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "main.layout", "welcome")
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("delivery/web/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", homeHandler)
	http.ListenAndServe(":2121", mux)
}

