package handler

import (
	"TrainProject/entity"
	"TrainProject/schedule/repository"
	"TrainProject/schedule/service"
	"TrainSystem/delivery/http/handler"

	"bytes"

	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestAdminSchedules(t *testing.T) {
	tmp1 := template.Must(template.ParseGlob("../../ui/templates/*"))

	scheduleRepo := repository.NewMockScheduleRepo(nil)
	scheduleServ := service.NewScheduleService(scheduleRepo)

	adminSchhHandler := handler.NewAdminScheduleHandler(tmp1, scheduleServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/schedules", adminSchhHandler.AdminSchedules)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/admin/schedules")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(body, []byte("Mock Schedule 01")) {
		t.Errorf("want body to contain %q", body)
	}
}

func TestAdminSchedulesNew(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	scheduleRepo := repository.NewMockScheduleRepo(nil)
	scheduleServ := service.NewScheduleService(scheduleRepo)

	adminSchtHandler := handler.NewAdminScheduleHandler(tmpl, scheduleServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/schedules/new", adminSchtHandler.AdminSchedulesNew)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}


	form.Add("TrainSource", entity.ScheduleMock.TrainSource)
	form.Add("TrainDestination", entity.ScheduleMock.TrainDestination)
	form.Add("Image", entity.ScheduleMock.Image)

	resp, err := tc.PostForm(sURL+"/admin/schedules/new", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Mock Schedule 01")) {
		t.Errorf("want body to contain %q", body)
	}

}




func TestAdminSchedulesUpdate(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	scheduleRepo := repository.NewMockScheduleRepo(nil)
	scheduleServ := service.NewScheduleService(scheduleRepo)

	adminSchtHandler := handler.NewAdminScheduleHandler(tmpl, scheduleServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/schedules/update", adminSchtHandler.AdminSchedules)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}

	form.Add("ID", string(entity.ScheduleMock.ID))
	form.Add("TrainSource", entity.ScheduleMock.TrainSource)
	form.Add("TrainDestination", entity.ScheduleMock.TrainDestination)
	form.Add("Image", entity.ScheduleMock.Image)

	resp, err := tc.PostForm(sURL+"/admin/schedules/update?id=1", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Mock Schedule 01")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestAdminSchedulesDelete(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	scheduleRepo := repository.NewMockScheduleRepo(nil)
	scheduleServ := service.NewScheduleService(scheduleRepo)

	adminCatgHandler := handler.NewAdminScheduleHandler(tmpl, scheduleServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/schedules/delete", adminCatgHandler.AdminSchedules)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}

	form.Add("ID", string(entity.ScheduleMock.ID))
	form.Add("TrainSource", entity.ScheduleMock.TrainSource)
	form.Add("TrainDestination", entity.ScheduleMock.TrainDestination)
	form.Add("Image", entity.ScheduleMock.Image)

	resp, err := tc.PostForm(sURL+"/admin/schedules/delete?id=1", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Mock Schedule 01")) {
		t.Errorf("want body to contain %q", body)
	}

}










































