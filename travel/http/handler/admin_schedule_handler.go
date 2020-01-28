package handler

import (
	"TrainProject/entity"
	"TrainProject/schedule"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type AdminScheduleHandler struct {
	tmp1 *template.Template
	scheduleSrv schedule.ScheduleService
}

func NewAdminScheduleHandler(t *template.Template, ss schedule.ScheduleService) *AdminScheduleHandler {
	return &AdminScheduleHandler{tmp1:t, scheduleSrv:ss}
}

func (ash *AdminScheduleHandler) AdminSchedules(w http.ResponseWriter, r *http.Request){
	schedules, errs := ash.scheduleSrv.Schedules()
	if errs != nil {
		panic(errs)
	}
	ash.tmp1.ExecuteTemplate(w, "admin.sched.layout", schedules)
}

func (ash *AdminScheduleHandler) AdminSchedulesNew(w http.ResponseWriter, r *http.Request){

	if r.Method == http.MethodPost {
		sch := &entity.Schedule{}
		sch.TrainSource = r.FormValue("trainSource")
		sch.TrainDestination= r.FormValue("trainDestination")

		mf, fh, err := r.FormFile("trainimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		sch.Image = fh.Filename

		writeFile(&mf, fh.Filename)

		_, errs := ash.scheduleSrv.StoreSchedule(sch)

		if len(errs) > 0 {
			panic(errs)
		}
		http.Redirect(w,r, "/admin/schedules", http.StatusSeeOther)
	} else {
		ash.tmp1.ExecuteTemplate(w, "admin.sched.new.layout", nil)
	}
}

func (ash *AdminScheduleHandler) AdminSchedulesUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method ==http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}
		sat, errs := ash.scheduleSrv.Schedule(uint(id))

		if len(errs) > 0 {
			panic(errs)
		}
		ash.tmp1.ExecuteTemplate(w, "admin.sched.update.layout", sat)
	}  else if r.Method == http.MethodPost {
		sch := &entity.Schedule{}
		id,_ := strconv.Atoi(r.FormValue("id"))
		sch.ID= uint(id)
		sch.TrainSource= r.FormValue("trainSource")
		sch.TrainDestination = r.FormValue("trainDestination")
		sch.Image = r.FormValue("image")

		mf, _, err := r.FormFile("trainimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()
		writeFile(&mf, sch.Image)

		_, errs := ash.scheduleSrv.UpdateSchedule(sch)
		if len(errs) > 0 {
			panic(errs)
		}
		http.Redirect(w,r, "/admin/schedules", http.StatusSeeOther)
	} else {
		http.Redirect(w,r, "/admin/schedules", http.StatusSeeOther)
	}
}

func (ash *AdminScheduleHandler) AdminSchedulesDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}
		_, errs := ash.scheduleSrv.DeleteSchedule(uint(id))
		if len(errs) > 0 {
			panic(err)
		}
	}
	http.Redirect(w,r, "/admin/schedules", http.StatusSeeOther)
}


func writeFile(mf *multipart.File, fname string)  {
	wd, err := os.Getwd()

	if err  != nil {
		panic(err)
	}

	path := filepath.Join(wd, "ui", "assets", "img", fname)
	image, err := os.Create(path)

	if err !=nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}


































