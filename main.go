package main

import (
	"TrainSystem/entity"
	"TrainSystem/train/repository"
	"TrainSystem/train/service"
	"database/sql"
	"strconv"

	"net/http"

)


func adminSchedulesDelete(w http.ResponseWriter, r *http.Request){
	var scheduleService *service.ScheduleService

	if r.Method == http.MethodGet{

		idRaw := r.URL.Query().Get("id")
		id, _ :=strconv.Atoi(idRaw)

		scheduleService.DeleteSchedule(id)
	}
	http.Redirect(w,r, "/admin/schedules", http.StatusSeeOther)
}
func adminSchedulesNew(w http.ResponseWriter, r * http.Request){
	if r.Method == http.MethodPost{

		sch := entity.Schedule{}
		sch.TrainSource = r.FormValue("TrainSource")
		sch.TrainDestination= r.FormValue("TrainDestination")

		scheduleService.StoreCatagory(sch)
	}
		http.Redirect(w,r, "/admin/schedules", http.StatusSeeOther)
}

func main() {
	dbconn, err := sql.Open("postgres", "postgres://app_admin:P@$$wordD2@localhost/traindb")

	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	schRepo := repository.NewScheduleRepositoryImp1(dbconn)
	schServ := service.NewScheduleService(schRepo)

	http.HandleFunc("/admin/schedules/delete", adminSchedulesDelete)
	http.HandleFunc("admin/schedules/new", adminSchedulesNew)
}




