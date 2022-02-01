package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Hangman struct {
	WordToGuess  string
	HiddenWord   []string
	UserInput    string
	Lives        int
	Proposition  []string
	FoundLetters int
	Win          bool
	Loose        bool
	File         string
}

var array []string

func main() {
	easy := Hangman{Lives: 10, Win: false, Loose: false, File: "wordsEasy.txt"}
	easy.hangmanInit()
	classic := Hangman{Lives: 10, Win: false, Loose: false, File: "words.txt"}
	classic.hangmanInit()
	fmt.Println("server starting")
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index", "")
	})
	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		classic.start(r)
		tmpl.ExecuteTemplate(w, "hangman", classic)
	})
	http.HandleFunc("/hangmanEasy", func(w http.ResponseWriter, r *http.Request) {
		easy.start(r)
		tmpl.ExecuteTemplate(w, "hangmanEasy", easy)
	})
	http.ListenAndServe(":80", nil)
	fmt.Println("server closing")
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
