package main

import (
	"html/template"
	"net/http"
)

type noah struct {
	Name    string
	Surname string
	Age     string
}

func User3() {
	tmpUser3 := template.Must(template.ParseFiles("templates/user3.html"))
	http.HandleFunc("/user3", func(w http.ResponseWriter, r *http.Request) {
		noahdata := noah{
			Name:    "Noah",
			Surname: "Bennet",
			Age:     "22",
		}
		tmpUser3.Execute(w, noahdata)
	})
}
