package main

import (
	"math/rand"
	"net/http"
	"time"
)

func (user *Hangman) start(r *http.Request) { //Programme de lancement du jeu
	verifLettersUsed := 0
	verifGoodProposition := 0
	user.UserInput = r.FormValue("userinput")
	for i := range user.Proposition {
		if user.UserInput == user.Proposition[i] {
			verifLettersUsed++
		}
	}
	if verifLettersUsed == 0 { //Ajouts aux propositions passées
		user.Proposition = append(user.Proposition, user.UserInput)
	}
	for i := 0; i < len(user.WordToGuess); i++ { //Vérification si la lettre est présente dans le mot
		if user.UserInput == string(user.WordToGuess[i]) && string(user.HiddenWord[i]) == "_" {
			user.HiddenWord[i] = string(user.WordToGuess[i])
			user.FoundLetters++
		} else {
			verifGoodProposition++
		}
	}
	if user.UserInput == user.WordToGuess { //Vérification si le mot a été trouvé (via une proposition de mot)
		user.Win = true
	}
	if user.FoundLetters == len(user.WordToGuess) {
		user.Win = true
	}
	if verifGoodProposition == len(user.WordToGuess) { //Modification du compteur d'essai en cas d'échec
		if len(user.UserInput) == 1 {
			user.Lives--
		} else if len(user.UserInput) > 1 {
			user.Lives -= 2
			if user.Lives < 0 {
				user.Lives = 0
			}
		} else {
			println("rien ne se passe...")
		}
		println("Not present in the word, ", user.Lives, " attempts remaining")
	}
	if user.Lives <= 0 {
		user.Loose = true
	}
	println()
}

func hideToFindWord(word string) []string { //Programme pour créer le mot caché
	var hiddenWord []string
	for i := 0; i < len(word); i++ {
		hiddenWord = append(hiddenWord, "_")
	}
	return hiddenWord
}

func (user *Hangman) hangmanInit() {
	fileScanner := createScanner("words.txt")
	array = getWords(fileScanner, array)
	rand.Seed(time.Now().UnixNano()) //Initialisation de l'aléatoire
	ran := rand.Intn(len(array))
	user.WordToGuess = array[ran]
	user.HiddenWord = hideToFindWord(user.WordToGuess)
}
