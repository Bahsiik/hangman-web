package main

import (
	"math/rand"
	"net/http"
)

func (user *Hangman) start(r *http.Request) { //Programme de lancement du jeu
	verifLettersUsed := 0
	verifGoodProposition := 0
	user.UserInput = r.FormValue("userinput")
	verifLettersUsed = alreadyUsed(user, verifLettersUsed)
	addProp(verifLettersUsed, user)
	verifGoodProposition = isPropTrue(user, verifGoodProposition)
	Win(user)
	livesChange(verifGoodProposition, user)
	Loose(user)
	println()
}

func alreadyUsed(user *Hangman, verifLettersUsed int) int {
	for i := range user.Proposition {
		if user.UserInput == user.Proposition[i] {
			verifLettersUsed++
		}
	}
	return verifLettersUsed
}

func addProp(verifLettersUsed int, user *Hangman) {
	if verifLettersUsed == 0 { //Ajouts aux propositions passées
		user.Proposition = append(user.Proposition, user.UserInput)
	}
}

func livesChange(verifGoodProposition int, user *Hangman) {
	if verifGoodProposition == len(user.WordToGuess) { //Modification du compteur d'essai en cas d'échec
		if len(user.UserInput) == 1 {
			user.Lives--
		} else if len(user.UserInput) > 1 {
			user.Lives -= 2
			if user.Lives < 0 {
				user.Lives = 0
			}
		} else {
		}
	}
}

func Loose(user *Hangman) {
	if user.Lives <= 0 {
		user.Loose = true
	}
}

func Win(user *Hangman) {
	if user.UserInput == user.WordToGuess { //Vérification si le mot a été trouvé (via une proposition de mot)
		user.Win = true
	}
	if user.FoundLetters == len(user.WordToGuess) {
		user.Win = true
	}
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

func isPropTrue(user *Hangman, verifGoodProposition int) int {
	for i := 0; i < len(user.WordToGuess); i++ { //Vérification si la lettre est présente dans le mot
		if user.UserInput == string(user.WordToGuess[i]) && string(user.HiddenWord[i]) == "_" {
			user.HiddenWord[i] = string(user.WordToGuess[i])
			user.FoundLetters++
		} else {
			verifGoodProposition++
		}
	}
	return verifGoodProposition
}
