package handler

import (
	"TrainSystem/train"
	"html/template"
	"net/http"
)

type TrainHandler struct {
	temp1 *template.Template
	scheduleSrv train.ScheduleService
}
func NewTrainHandler(T *template.Template, CS train.ScheduleService) *TrainHandler{
	return &TrainHandler{temp1:T, scheduleSrv:CS}
}

func (th *TrainHandler) Index(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	schedules, err := th.scheduleSrv.Schedules()
	if err != nil {
		panic(err)
	}
	th.temp1.ExecuteTemplate(w, "index.layout", schedules)

}

func (th *TrainHandler) Train (w http.ResponseWriter, r *http.Request){
	th.temp1.ExecuteTemplate(w, "train.layout", nil)
}
























