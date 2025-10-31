package main

import (
	"html/template"
	"net/http"
)

type Jackson struct {
	Name    string
	Surname string
	Age     string
}

func User8() {
	tmp := template.Must(template.ParseFiles("templates/user8.html"))
	http.HandleFunc("/user8", func(w http.ResponseWriter, r *http.Request) {
		jacksondata := Jackson{
			Name:    "Sophia",
			Surname: "Turner",
			Age:     "39",
		}
		tmp.Execute(w, jacksondata)
	})
}
