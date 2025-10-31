package main

import (
	"html/template"
	"net/http"
)

type eth struct {
	Name    string
	Surname string
	Age     string
}

func User1() {
	tmp := template.Must(template.ParseFiles("templates/user1.html"))
	http.HandleFunc("/user1", func(w http.ResponseWriter, r *http.Request) {
		ethdata := eth{
			Name:    "Ethan",
			Surname: "Carter",
			Age:     "30",
		}
		tmp.Execute(w, ethdata)
	})
}
