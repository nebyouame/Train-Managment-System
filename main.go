package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/betsegawlemma/web-prog-lab-05-mem/entity"
	"github.com/betsegawlemma/web-prog-lab-05-mem/menu/service"
)

var tmpl = template.Must(template.ParseGlob("delivery/web/templates/*"))
var ScheduleCache service.ScheduleCache

func index(w http.ResponseWriter, r *http.Request) {

	idRaw := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		id = 1
	}

	cat, err := ScheduleCache.Schedule(id)
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(w, "index.layout", cat)
}
func book(w http.ResponseWriter, r *http.Request) {

	idRaw := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		id = 1
	}

	cat, err := ScheduleCache.Schedule(id)
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(w, "index.layout", cat)
}

func init() {
	ScheduleCache = service.NewScheduleCache()
	Train1 := entity.Schedule{ID: 1, Name: "First train", Description: "From Station1 to Station3", Image: "sss.png"}
	Train2 := entity.Schedule{ID: 2, Name: "Second train", Description: "From Station2 to Station3", Image: "lnc.png"}
	Train3 := entity.Schedule{ID: 3, Name: "Third train", Description: "From Station1 to Station2", Image: "dnr.png"}

	ScheduleCache.StoreSchedule(&Train1)
	ScheduleCache.StoreSchedule(&Train2)
	ScheduleCache.StoreSchedule(&Train3)

}

func main() {
	fs := http.FileServer(http.Dir("delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8181", nil)
}
