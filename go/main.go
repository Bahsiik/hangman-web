package main

import (
	"html/template"
	"net/http"
)

func main() {
	check := 0
	tmpl := template.Must(template.ParseGlob("templates/*.gohtml"))
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	images := http.FileServer(http.Dir("images"))
	http.Handle("/images/", http.StripPrefix("/images/", images))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		check = 0
		easy = Hangman{Lives: 10, Win: false, Loose: false, File: "./text/easy.txt"}
		normal = Hangman{Lives: 10, Win: false, Loose: false, File: "./text/normal.txt"}
		hard = Hangman{Lives: 10, Win: false, Loose: false, File: "./text/hard.txt"}
		tmpl.ExecuteTemplate(w, "index", "")
	})
	http.HandleFunc("/hangmanEasy", func(w http.ResponseWriter, r *http.Request) {
		for check == 0 {
			easy.getRandomWord()
			check += 1
		}
		easy.start(r)
		err := tmpl.ExecuteTemplate(w, "hangmanEasy", easy)
		if err != nil {
			return
		}
	})
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		for check == 0 {
			normal.getRandomWord()
			check += 1
		}
		normal.start(r)
		err := tmpl.ExecuteTemplate(w, "hangman", normal)
		if err != nil {
			return
		}
	})
	http.HandleFunc("/hangmanHard", func(w http.ResponseWriter, r *http.Request) {
		for check == 0 {
			hard.getRandomWord()
			check += 1
		}
		hard.start(r)
		err := tmpl.ExecuteTemplate(w, "hangmanHard", hard)
		if err != nil {
			return
		}
	})
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}
