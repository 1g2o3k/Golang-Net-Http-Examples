package main

import (
	"html/template"
	"net/http"
)

type chl struct {
	Name    string
	Surname string
	Age     string
}

func User6() {
	tmp := template.Must(template.ParseFiles("templates/user6.html"))
	http.HandleFunc("/user6", func(w http.ResponseWriter, r *http.Request) {
		chldata := chl{
			Name:    "Chloe",
			Surname: "Harrison",
			Age:     "45",
		}
		tmp.Execute(w, chldata)
	})
}
