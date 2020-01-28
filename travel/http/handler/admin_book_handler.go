package handler

import (
	"TrainProject/entity"
	"TrainProject/schedule"
	"html/template"
	"net/http"
	"strconv"
)

type AdminBookHandler struct {
	tmp1 *template.Template
	bookSrv schedule.BookService
}

func NewAdminBookHandler(t *template.Template, bs schedule.BookService) *AdminBookHandler {
	return &AdminBookHandler{tmp1:t, bookSrv:bs}
}

func (abh *AdminBookHandler) AdminBooks(w http.ResponseWriter, r *http.Request){
	books, errs := abh.bookSrv.Books()
	if errs != nil {
		panic(errs)
	}
	abh.tmp1.ExecuteTemplate(w, "admin.books.layout", books)

}

func (abh *AdminBookHandler) AdminBooksNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		bkk := &entity.Book{}
		bkk.UserID = r.FormValue("userID")
		bkk.InfoID = r.FormValue("infoID")

		mf, fh, err := r.FormFile("bookimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		bkk.Image = fh.Filename


		_, errs := abh.bookSrv.StoreBook(bkk)
		if len(errs) > 0 {
			panic(errs)
		}
		http.Redirect(w,r, "/admin/books", http.StatusSeeOther)
	}	else {
		abh.tmp1.ExecuteTemplate(w, "admin.books.new.layout", nil)
	}
}

func (abh *AdminBookHandler) AdminBooksDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)
		if err != nil {
			panic(err)
		}
		_, errs := abh.bookSrv.DeleteBook(uint(id))
		if len(errs) > 0 {
			panic(err)
		}
		}
		http.Redirect(w,r, "/admin/books", http.StatusSeeOther)
}

func (abh *AdminBookHandler) AdminBooksUpdate(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}
		mat, errs := abh.bookSrv.Book(uint(id))
		if len(errs) > 0 {
			panic(errs)
		}
		abh.tmp1.ExecuteTemplate(w, "admin.books.update.layout", mat)

	}	else if r.Method==http.MethodPost {
		bkk := &entity.Book{}
		id,_ := strconv.Atoi(r.FormValue("id"))
		bkk.ID= uint(id)
		bkk.UserID = r.FormValue("userID")
		bkk.InfoID = r.FormValue("infoID")
		bkk.Image= r.FormValue("image")

		mf, _, err := r.FormFile("bookimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()
		writeFile(&mf, bkk.Image)

		_, errs := abh.bookSrv.UpdateBook(bkk)
		if len(errs) > 0 {
			panic(errs)
		}
		http.Redirect(w,r, "/admin/books", http.StatusSeeOther)
	}	else {
			http.Redirect(w,r, "/admin/books", http.StatusSeeOther)
	}
}
























































