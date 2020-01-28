package main

import (
	"TrainProject/entity"
	"TrainProject/schedule/repository"
	"TrainProject/schedule/service"
	"TrainProject/travel/http/handler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"html/template"
	"net/http"
)

func createTable(dbconn *gorm.DB) []error {
	errs := dbconn.CreateTable(&entity.Info{}, &entity.Book{}, &entity.Schedule{}, &entity.User{}, &entity.Role{}, &entity.Path{}, &entity.Comment{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}

func main() {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:Ermi12345@localhost/traindb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	createTable(dbconn)

	tmp1 := template.Must(template.ParseGlob("ui/templates/*"))

	scheduleRepo := repository.NewScheduleGormRepo(dbconn)
	scheduleServ := service.NewScheduleService(scheduleRepo)

	

	roleRepo := repository.NewRoleGormRepo(dbconn)
	roleServ := service.NewRoleService(roleRepo)

	

	adminSatHandler := handler.NewAdminScheduleHandler(tmp1, scheduleServ)
	scheduleHandler := handler.NewScheduleHandler(tmp1, scheduleServ)

	

	adminRoleHandler := handler.NewAdminRoleHandler(tmp1, roleServ)



	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", scheduleHandler.Index)
	http.HandleFunc("/schedule", scheduleHandler.Schedule)
	http.HandleFunc("/admin", scheduleHandler.Admin)


	http.HandleFunc("/admin/schedules", adminSatHandler.AdminSchedules)
	http.HandleFunc("/admin/schedules/new", adminSatHandler.AdminSchedulesNew)
	http.HandleFunc("/admin/schedules/update", adminSatHandler.AdminSchedulesUpdate)
	http.HandleFunc("/admin/schedules/delete", adminSatHandler.AdminSchedulesDelete)



	http.HandleFunc("/admin/roles", adminRoleHandler.AdminRoles)
	http.HandleFunc("/admin/roles/new", adminRoleHandler.AdminRolesNew)
	http.HandleFunc("/admin/roles/update", adminRoleHandler.AdminRolesUpdate)
	http.HandleFunc("/admin/roles/delete", adminRoleHandler.AdminRolesDelete)

	


	http.ListenAndServe(":8181", nil)

}