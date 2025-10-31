package main

import (
	"html/template"
	"net/http"
)

type lucas struct {
	Name    string
	Surname string
	Age     string
}

func User7() {
	tmp := template.Must(template.ParseFiles("templates/user6.html"))
	http.HandleFunc("/user7", func(w http.ResponseWriter, r *http.Request) {
		lucasdata := lucas{
			Name:    "Lucas",
			Surname: "Morgan",
			Age:     "38",
		}
		tmp.Execute(w, lucasdata)
	})
}
