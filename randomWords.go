package main

import (
	"math/rand"
	"time"
)

func hideToFindWord(word string) []string { //Programme pour créer le mot caché
	var hiddenWord []string
	for i := 0; i < len(word); i++ {
		hiddenWord = append(hiddenWord, "_")
	}
	return hiddenWord
}

func (user *Hangman) hangmanInit() {
	fileScanner := createScanner(user.Files)
	array = getWords(fileScanner, array)
	rand.Seed(time.Now().UnixNano()) //Initialisation de l'aléatoire
	ran := rand.Intn(len(array))
	user.WordToGuess = array[ran]
	user.HiddenWord = hideToFindWord(user.WordToGuess)
}
