package main

import (
	"TrainSystem/delivery/http/handler"
	"TrainSystem/entity"
	"TrainSystem/train/repository"
	"TrainSystem/train/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"html/template"
	"net/http"
)

func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.CreateTable(&entity.Item{}, &entity.Schedule{}, &entity.Role{}, &entity.User{})
	if errs !=nil{
		panic(errs)
	}
	return nil
}
func main(){
	dbconn, err := gorm.Open("postgres", "postgres://postgres:Ermi12345@localhost/traindb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	templ := template.Must(template.ParseGlob("TrainSystem/ui/templates/*"))

	serviceRepo := repository.NewScheduleGormRepo(dbconn)
	serviceServ := service.NewScheduleService(serviceRepo)

	adminSchgHandler := handler.NewAdminScheduleHandler(templ, serviceServ)
	trainHandler := handler.NewTrainHandler(templ, serviceServ)


	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", trainHandler.Index)
	http.HandleFunc("/train", trainHandler.Train)
	http.HandleFunc("/admin/schedules", adminSchgHandler.AdminSchedules)
	http.HandleFunc("/admin/schedules/new", adminSchgHandler.AdminSchedulesNew)
	http.HandleFunc("/admin/schedules/update", adminSchgHandler.AdminSchedulesUpdate)
	http.HandleFunc("/admin/schedules/delete", adminSchgHandler.AdminSchedulesDelete)

	http.ListenAndServe(":8181", nil)
}
