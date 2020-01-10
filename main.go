package main

import (
	"TrainSystem/entity"
	"TrainSystem/menu/service"
	"html/template"
	"net/http"
	"strconv"
)

var templ = template.Must(template.ParseGlob("delivery/web/templates/*"))
var scheduleCache service.ScheduleCache

func index(w http.ResponseWriter, r *http.Request) {

	idRaw := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		id = 1
	}
	cat, err := scheduleCache.Schedule(id)
	if err != nil {
		panic(err)
	}
	templ.ExecuteTemplate(w, "index.layout", cat)
}

func init(){
	scheduleCache = service.NewSchduleCache()
	Train1 := entity.Schedule{ID: 1, TrainSource: "From Source1", TrainDestination: "To destination2", Image: "train.png"}
	Train2 := entity.Schedule{ID: 2, TrainSource: "From Source2", TrainDestination: "To destination3", Image: "ggg.png"}
	Train3 := entity.Schedule{ID: 3, TrainSource: "From Source3", TrainDestination: "To destination1", Image: "ggg.png"}
	scheduleCache.StoreSchedule(&Train1)
	scheduleCache.StoreSchedule(&Train2)
	scheduleCache.StoreSchedule(&Train3)
}

func main() {
	fs := http.FileServer(http.Dir("delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	
	http.ListenAndServe(":8181", nil)
}


