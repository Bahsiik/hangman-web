package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Pseudo string
}

func main() {
	fmt.Println("server starting")

	tmpl, _ := template.ParseGlob("templates/*.html")

	fs := http.FileServer(http.Dir("../css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := User{
			Pseudo: r.FormValue("pseudo"),
		}
		err := tmpl.ExecuteTemplate(w, "index", user)
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
	fmt.Println("server closing")
}
