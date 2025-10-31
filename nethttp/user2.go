package main

import (
	"html/template"
	"net/http"
)

type lily struct {
	Name    string
	Surname string
	Age     string
}

func User2() {
	tmpUser2 := template.Must(template.ParseFiles("templates/user2.html"))
	http.HandleFunc("/user2", func(w http.ResponseWriter, r *http.Request) {
		lilydata := lily{
			Name:    "Lily",
			Surname: "Thompson",
			Age:     "20",
		}
		tmpUser2.Execute(w, lilydata)
	})
}
