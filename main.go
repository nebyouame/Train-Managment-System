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

func main() {
	dbconn, err := sql.Open("postgres", "postgres://app_admin:P@$$wordD2@localhost/traindb")

	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	schRepo := repository.NewScheduleRepositoryImp1(dbconn)
	schServ := service.NewScheduleService(schRepo)

	http.HandleFunc("/admin/schedules/delete", adminSchedulesDelete)
}




