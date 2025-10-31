package main

import (
	"html/template"
	"net/http"
)

type oli struct {
	Name    string
	Surname string
	Age     string
}

func User5() {
	tmp5 := template.Must(template.ParseFiles("templates/user5.html"))
	http.HandleFunc("/user5", func(w http.ResponseWriter, r *http.Request) {
		olidata := oli{
			Name:    "Oliver",
			Surname: "Reed",
			Age:     "18",
		}
		tmp5.Execute(w, olidata)
	})

}
