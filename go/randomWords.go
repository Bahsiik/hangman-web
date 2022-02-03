package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func (user *Hangman) getRandomWord() { //Choix du mot aléatoirement dans le dossier .txt
	var array []string
	fileScanner := createScanner(user.File)
	array = getWords(fileScanner, array)
	rand.Seed(time.Now().UnixNano()) //Initialisation de l'aléatoire
	ran := rand.Intn(len(array))
	user.WordToGuess = array[ran]
	user.HiddenWord = hideToFindWord(user.WordToGuess)
	user.FoundLetters = user.showToFindLetters()
}

func hideToFindWord(word string) []string {
	var hiddenWord []string
	for i := 0; i < len(word); i++ {
		hiddenWord = append(hiddenWord, "_")
	}
	return hiddenWord
}

func (user *Hangman) showToFindLetters() int { //Choix des lettres affichées dès le début
	lettersToDisplay := (len(user.HiddenWord) / 2) - 1
	var displayedLetters int
	for i := 0; i < lettersToDisplay; i++ {
		index := rand.Intn(len(user.HiddenWord))
		if user.HiddenWord[index] == "_" {
			displayedLetters++
		}
		user.HiddenWord[index] = string(user.WordToGuess[index])
	}
	return displayedLetters
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
