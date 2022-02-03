package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	easy := Hangman{Lives: 10, Win: false, Loose: false, File: "wordsEasy.txt"}
	easy.getRandomWord()
	hard := Hangman{Lives: 10, Win: false, Loose: false, File: "wordsHard.txt"}
	hard.getRandomWord()
	classic := Hangman{Lives: 10, Win: false, Loose: false, File: "words.txt"}
	classic.getRandomWord()

	fmt.Println("server starting")
	tmpl := template.Must(template.ParseGlob("templates/*.gohtml"))
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	images := http.FileServer(http.Dir("images"))
	http.Handle("/images/", http.StripPrefix("/images/", images))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index", "")
	})
	http.HandleFunc("/hangmanEasy", func(w http.ResponseWriter, r *http.Request) {
		easy.start(r)
		tmpl.ExecuteTemplate(w, "hangmanEasy", easy)
	})
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		classic.start(r)
		tmpl.ExecuteTemplate(w, "hangman", classic)
	})
	http.HandleFunc("/hangmanHard", func(w http.ResponseWriter, r *http.Request) {
		hard.start(r)
		tmpl.ExecuteTemplate(w, "hangmanHard", hard)
	})
	http.ListenAndServe(":80", nil)
}

func getWords(fileScanner *bufio.Scanner, array []string) []string { //Programme de récupération des mots du fichier txt
	for fileScanner.Scan() {
		array = append(array, fileScanner.Text())
	}
	return array
}

func createScanner(fileName string) *bufio.Scanner { //Programme de création d'un scanner
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	return fileScanner
}
