package main

import (
	"math/rand"
	"time"
)

var array []string

func (user *Hangman) getRandomWord() { //Choix du mot aléatoirement dans le dossier .txt
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
