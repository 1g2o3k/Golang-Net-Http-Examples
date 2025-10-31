package main

import (
	"html/template"
	"net/http"
)

type nathan struct{
	Name string
	Surname string
	Age string
}
func User12(){
tmp12:=template.Must(template.ParseFiles("templates/user12.html"))
http.HandleFunc("/user12",func (w http.ResponseWriter, r *http.Request) {
	nathandata:=nathan{
		Name: "Nathan",
		Surname: "Corren",
		Age: "32",
	}
	tmp12.Execute(w,nathandata)
})	

}