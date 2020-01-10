package main

import (
	"html/template"
	"net/http"

)

var templ = template.Must(template.ParseGlob("delivery/web/templates/*"))

func index(w http.ResponseWriter, r *http.Request){
	templ.ExecuteTemplate(w, "index.layout", nil)
}
func Schedule(w http.ResponseWriter, r *http.Request){
	templ.ExecuteTemplate(w, "Schedule.layout", nil)
}

func main() {
	fs := http.FileServer(http.Dir("delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/Schedule", Schedule)
	http.ListenAndServe(":8181", nil)
}


