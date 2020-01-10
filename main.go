package main

import (
	"TrainSystem/entity"
	"TrainSystem/menu/service"

	"html/template"
	"net/http"

)

var templ = template.Must(template.ParseGlob("delivery/web/templates/*"))
var scheduleService *service.ScheduleService

func index(w http.ResponseWriter, r *http.Request) {

	Schedule, err := scheduleService.Schedules()
	if err != nil {
		panic(err)
	}
	templ.ExecuteTemplate(w, "index.layout", Schedule)
}


func init(){
	scheduleService = service.NewScheduleService("catagory.csv")

	schedules := []entity.Schedule{
		entity.Schedule{ID: 1, TrainSource: "From s1", TrainDestination: "TO D2", Image: "train.png"},
		entity.Schedule{ID: 2, TrainSource: "From s2", TrainDestination: "TO D3", Image: "train.png"},
		}
	scheduleService.StoreSchedules(schedules)
}

func main() {
	fs := http.FileServer(http.Dir("delivery/web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	
	http.ListenAndServe(":8181", nil)
}


