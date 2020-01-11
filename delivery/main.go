package delivery

import (
	"TrainSystem/delivery/http/handler"
	"TrainSystem/entity"
	"TrainSystem/train/repository"
	"TrainSystem/train/service"
	"github.com/jinzhu/gorm"
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

	templ := template.Must(template.ParseGlob("ui/templates.*"))
	serviceRepo := repository.NewScheduleGormRepo(dbconn)
	serviceServ := service.NewScheduleService(serviceRepo)

	trainHandler := handler.NewTrainHandler(templ, serviceServ)


	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", trainHandler.Index)
	http.HandleFunc("/train", trainHandler.Train)


	http.ListenAndServe(":8181", nil)
}
