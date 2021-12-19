package main

import (
	"fmt"
	"net/http"

	"encoding/json"
	"io/ioutil"
	"html/template"

)

func Home(w http.ResponseWriter, r *http.Request){

	t:= template.Must(template.ParseFiles("ui/pages/home.html"))
	t.Execute(w,nil)
}
func About(w http.ResponseWriter, r *http.Request){

	t:= template.Must(template.ParseFiles("ui/pages/about.html"))
	t.Execute(w,nil)
}
func Projects(w http.ResponseWriter, r *http.Request){

	t:= template.Must(template.ParseFiles("ui/pages/loyihalar.html"))
	t.Execute(w,nil)
}
func Techno(w http.ResponseWriter, r *http.Request){

	t:= template.Must(template.ParseFiles("ui/pages/techno.html"))
	t.Execute(w,nil)
}
func Send(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Access-Control-Allow-Origin","*")
	var redirectTarget string

	if r.Method == "POST" {

		var bodyData CommentInfo
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		
		json.Unmarshal(body, &bodyData)

		fmt.Println(bodyData)
		
		err = SendLink(bodyData)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
		redirectTarget = "/about"
	}
	http.Redirect(w, r, redirectTarget, 302)
	
}
