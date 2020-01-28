package handler

import (
	"TrainProject/schedule"
	"html/template"
	"net/http"
)

type ScheduleHandler struct {
	tmp1 *template.Template
	scheduleSrv schedule.ScheduleService
}

func NewScheduleHandler(T *template.Template, CS schedule.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{tmp1:T, scheduleSrv:CS}
}

func (sh *ScheduleHandler) Index(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path !="/" {
		http.NotFound(w, r)
		return
	}
	schedules, err := sh.scheduleSrv.Schedules()
	if err != nil {
		panic(err)
	}
	sh.tmp1.ExecuteTemplate(w, "index.layout", schedules)
}

func (sh *ScheduleHandler) Schedule(w http.ResponseWriter, r *http.Request){
	sh.tmp1.ExecuteTemplate(w, "schedule.layout", nil)
}

func (sh *ScheduleHandler) Admin(w http.ResponseWriter, r *http.Request)  {
	sh.tmp1.ExecuteTemplate(w, "admin.index.layout", nil)
}





















