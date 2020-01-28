package handler

import (
	"TrainProject/entity"
	"TrainProject/schedule"
	"html/template"
	"net/http"
	"strconv"
)

type AdminUserHandler struct {
	tmp1 *template.Template
	userSrv schedule.UserService
}

func NewAdminUserHandler(t *template.Template, us schedule.UserService) *AdminUserHandler {
	return &AdminUserHandler{tmp1:t, userSrv:us}
}

func (auh *AdminUserHandler) AdminUsers(w http.ResponseWriter, r *http.Request){
	users, errs := auh.userSrv.Users()
	if errs != nil {
		panic(errs)
	}
	auh.tmp1.ExecuteTemplate(w, "admin.users.layout", users)
}
func (auh *AdminUserHandler) AdminUsersNew(w http.ResponseWriter, r *http.Request){

	if r.Method == http.MethodPost {
		sch := &entity.User{}

		sch.UserName = r.FormValue("userName")
		sch.FullName = r.FormValue("fullName")
		sch.Email= r.FormValue("email")
		sch.Phone= r.FormValue("phone")
		sch.Password = r.FormValue("password")

		_, errs := auh.userSrv.StoreUser(sch)

		if len(errs) > 0 {
			panic(errs)
		}
		http.Redirect(w,r, "/admin/users", http.StatusSeeOther)
	} else {
		auh.tmp1.ExecuteTemplate(w, "admin.users.new.layout", nil)
	}
}

func (auh *AdminUserHandler) AdminUsersUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method ==http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}
		sat, errs := auh.userSrv.User(uint(id))

		if len(errs) > 0 {
			panic(errs)
		}
		auh.tmp1.ExecuteTemplate(w, "admin.users.update.layout", sat)
	}  else if r.Method == http.MethodPost {
		sch := &entity.User{}
		id,_ := strconv.Atoi(r.FormValue("id"))
		sch.ID= uint(id)
		sch.UserName= r.FormValue("userName")
		sch.FullName= r.FormValue("fullName")
		sch.Email= r.FormValue("email")
		sch.Phone= r.FormValue("phone")
		sch.Password= r.FormValue("password")

		_, errs := auh.userSrv.UpdateUser(sch)
		if len(errs) > 0 {
			panic(errs)
		}
		http.Redirect(w,r, "/admin/users", http.StatusSeeOther)
	} else {
		http.Redirect(w,r, "/admin/users", http.StatusSeeOther)
	}
}

func (auh *AdminUserHandler) AdminUsersDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}
		_, errs := auh.userSrv.DeleteUser(uint(id))
		if len(errs) > 0 {
			panic(err)
		}
	}
	http.Redirect(w,r, "/admin/users", http.StatusSeeOther)
}


