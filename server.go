package main

import "net/http"

import "text/template"

var templ = template.Must(template.ParseGlob("delivery/web/templates/*"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "main.layout", "welcome")
}

func main() {
	mux := http.NewServeMux()

	//File server for bootstrap and css
	fs := http.FileServer(http.Dir("delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":2121", mux)

}
