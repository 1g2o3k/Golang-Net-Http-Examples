package main

import (
	"html/template"
    "net/http"
)
type adam struct{
	Name string
	Surname string
	Age string
}
func User9()  {
	tmp9:=template.Must(template.ParseFiles("templates/user9.html"))
    http.HandleFunc("/user9",func (w http.ResponseWriter, r *http.Request) {
		adamdata:=adam{
			Name: "Adam",
			Surname: "Smith",
			Age: "19",
		}
		tmp9.Execute(w,adamdata)
	})	
}