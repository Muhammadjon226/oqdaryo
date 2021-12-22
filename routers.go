package main

import (
	"fmt"
	"net/http"
	"strconv"

	"encoding/json"
	"html/template"

	"github.com/gorilla/mux"
	// "io/ioutil"
)

func Home(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("ui/pages/home.html"))
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		var bodyData CommentInfo = CommentInfo{
			Name:    r.FormValue("sender_name"),
			Email:   r.FormValue("sender_email"),
			Comment: r.FormValue("sender_comments"),
		}
		fmt.Println(bodyData)

		err := PostInfo(bodyData)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			http.Redirect(w, r, "/", 302)
		}
	}

}
func About(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("ui/pages/about.html"))
	t.Execute(w, nil)
}
func Projects(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("ui/pages/loyihalar.html"))
	t.Execute(w, nil)
}
func Techno(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("ui/pages/techno.html"))
	t.Execute(w, nil)
}

func Loggin(w http.ResponseWriter, r *http.Request) {

	redirectTarget := "/"

	if r.Method == "GET" {

		t := template.Must(template.ParseFiles("ui/pages/login.html"))
		t.Execute(w, nil)
	} else if r.Method == "POST" {

		name := r.FormValue("name")
		pass := r.FormValue("password")
		if name == "a" && pass == "a" {

			CreateToken(name, w)
			redirectTarget = "/admin/comments"
		}
	}

	http.Redirect(w, r, redirectTarget, 302)
}

func GetAll(w http.ResponseWriter, r *http.Request) {

	datas, err := GetAllInfo()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("ui/pages/admin.html"))
		t.Execute(w, nil)

	} else if r.Method == "VIEW" {

		e := json.NewEncoder(w)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		e.Encode(datas)

	}
}
func GetById(w http.ResponseWriter, r *http.Request) {
	redirectTarget := "/"

	v := mux.Vars(r)
	index, err := strconv.Atoi(v["commentId"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("ui/pages/admininfo.html"))
		t.Execute(w, nil)

	} else if r.Method == "VIEW" {

		e := json.NewEncoder(w)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		datas, err := GetAllInfo()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			panic(err)
		}
		fmt.Println(datas[index-1])
		e.Encode(datas[index-1])
	}
	http.Redirect(w, r, redirectTarget, 302)
}
