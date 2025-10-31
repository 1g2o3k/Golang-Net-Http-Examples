package main

import (
	"html/template"
	"net/http"
)

type tessa struct {
	Name    string
	Surname string
	Age     string
}

func User10() {
	tmp10 := template.Must(template.ParseFiles("templates/user10.html"))
	http.HandleFunc("/user10", func(w http.ResponseWriter, r *http.Request) {
		tessadata := tessa{
			Name:    "Tessa",
			Surname: "Grindle",
			Age:     "21",
		}
		tmp10.Execute(w, tessadata)
	})
}
