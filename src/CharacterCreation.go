package adventure

import (
	"fmt"
	"strings"
	"unicode"
)

func CharacterCreation() *Character {
	var nom string
	var classe string

	// Choix du nom
	for {
		fmt.Print("Entrez le nom de votre personnage : ")
		fmt.Scanln(&nom)

		if isAlpha(nom) {
			nom = formatNom(nom)
			break
		} else {
			fmt.Println("Le nom ne doit contenir que des lettres.")
		}
	}

	// Choix de la classe
	for {
		fmt.Print("Choisissez votre classe (Humain, Elfe, Nain) : ")
		fmt.Scanln(&classe)
		classe = strings.Title(strings.ToLower(classe)) // Formatage propre

		if classe == "Humain" || classe == "Elfe" || classe == "Nain" {
			break
		} else {
			fmt.Println("Classe invalide. Choisissez parmi Humain, Elfe ou Nain.")
		}
	}

	// Création du personnage
	personnage := Init(nom, classe)
	return personnage
}

func Init(nom string, classe string) *Character {
	var pvMax int

	switch classe {
	case "Humain":
		pvMax = 100
	case "Elfe":
		pvMax = 80
	case "Nain":
		pvMax = 120
	default:
		pvMax = 100 // Valeur par défaut si classe inconnue
	}

	return InitCharacter(nom, classe, 4, pvMax, pvMax/2, []Items{
		InitItems("potion", 3),
		InitItems("épée", 2),
	}, []string{"coup de poing"})
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func formatNom(nom string) string {
	nom = strings.ToLower(nom)
	return strings.Title(nom)
}
