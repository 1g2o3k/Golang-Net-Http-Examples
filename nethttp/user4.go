package main

import (
	"html/template"
	"net/http"
)

type ava struct {
	Name    string
	Surname string
	Age     string
}

func User4() {
	tmpUser4 := template.Must(template.ParseFiles("templates/user4.html"))
	http.HandleFunc("/user4", func(w http.ResponseWriter, r *http.Request) {
		avadata := ava{
			Name:    "Ava",
			Surname: "Mitchel",
			Age:     "35",
		}
		tmpUser4.Execute(w, avadata)
	})
}
