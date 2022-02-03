package main

import (
	"net/http"
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

func alreadyUsed(user *Hangman, verifLettersUsed int) int { // Fonction pour vérifier si la proposition a déjà été faite
	for i := range user.Proposition {
		if user.UserInput == user.Proposition[i] {
			verifLettersUsed++
		}
	}
	return verifLettersUsed
}

func addProp(verifLettersUsed int, user *Hangman) { //Ajouts de la proposition à la liste des propositions
	if verifLettersUsed == 0 {
		user.Proposition = append(user.Proposition, user.UserInput)
	}
}

func livesChange(verifGoodProposition int, user *Hangman) { //Modification du compteur d'essai en cas d'échec
	if verifGoodProposition == len(user.WordToGuess) {
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

func Loose(user *Hangman) { // Fonction pour activer la défaite
	if user.Lives <= 0 {
		user.Loose = true
	}
}

func Win(user *Hangman) { //Vérification si le mot a été trouvé (via une proposition de mot)
	if user.UserInput == user.WordToGuess {
		user.Win = true
	}
	if user.FoundLetters == len(user.WordToGuess) {
		user.Win = true
	}
}

func isPropTrue(user *Hangman, verifGoodProposition int) int { //Vérification si la lettre proposée est présente dans le mot
	for i := 0; i < len(user.WordToGuess); i++ {
		if user.UserInput == string(user.WordToGuess[i]) && string(user.HiddenWord[i]) == "_" {
			user.HiddenWord[i] = string(user.WordToGuess[i])
			user.FoundLetters++
		} else {
			verifGoodProposition++
		}
	}
	return verifGoodProposition
}
