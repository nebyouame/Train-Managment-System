package main

import (
	"html/template"
	"net/http"

)

var templ = template.Must(template.ParseGlob("delivery/web/templates/*"))

func index(w http.ResponseWriter, r *http.Request){
	templ.ExecuteTemplate(w, "index.layout", nil)
}


func main() {
	fs := http.FileServer(http.Dir("delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)

	http.ListenAndServe(":8181", nil)
}


