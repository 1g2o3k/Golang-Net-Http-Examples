package main

import (
	"html/template"
	"net/http"
)
type clara struct{
	Name string
	Surname string
	Age string	
}
func User11(){
	tmp11:=template.Must(template.ParseFiles("templates/user11.html"))
    http.HandleFunc("/user11",func (w http.ResponseWriter, r *http.Request) {
		claradata:=clara{
			Name: "Clara",
			Surname: "Winsley",
			Age: "35",
		}
		tmp11.Execute(w,claradata)
	})
}