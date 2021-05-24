package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)
	http.HandleFunc("/req", req)
	http.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":5555", nil)
}

// home page
func home(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	temp.Execute(w, nil)
}

// features page
func features(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	feat, err := temp.ParseFiles("wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	feat.Execute(w, nil)
}

// docs page
func docs(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	doc, err := temp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	doc.Execute(w, nil)
}

// form submit route method -1
// func req(w http.ResponseWriter, r *http.Request) {

// 	name := r.FormValue("name")
// 	company := r.FormValue("company")
// 	email := r.FormValue("email")

// 	fmt.Println(name, company, email)
// 	fmt.Fprintf(w, "Successfully submited your form ")
//  }

// .............method -2 advance method....
func req(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	for key, val := range r.Form {
		fmt.Println(key, val)
		fmt.Fprintf(w, `Submit Successfull`)
	}
}
