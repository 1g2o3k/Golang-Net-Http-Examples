package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmpindex := template.Must(template.ParseFiles("templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpindex.Execute(w, nil)
	})
	User1()
	User2()
	User3()
	User4()
	User5()
	User6()
	User7()
	User8()
	User9()
	User10()
	User11()
	User12()
	http.ListenAndServe(":8080", nil)
}
