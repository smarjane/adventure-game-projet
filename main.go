package main

import (
	adventure "adventure/src"
	"fmt"
	"os"
	"strings"
)

func main() {
	// tache 1
	c1 := adventure.CharacterCreation()
	adventure.TakePot(c1)
	adventure.SpellBook(c1)
	nouvelItem := adventure.InitItems("lampe torche", 2)
	adventure.ItemsMax(c1, nouvelItem)
	if adventure.ItemsMax(c1, nouvelItem) {
		c1.Inventaire = append(c1.Inventaire, nouvelItem)
	}

	// Tâche 6
	for {
		fmt.Println("\n--- Menu Principal ---")
		fmt.Println("1: Info de Character")
		fmt.Println("2: L'inventaire")
		fmt.Println("3: Quitter")
		// ajouter apres

		var choix string
		fmt.Println("Entrer votre choix")
		fmt.Scanln(&choix)
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			adventure.DisplayInfo(c1)
		case "2":
			adventure.AccessInventaire(c1)
		case "3":
			fmt.Println("En cours quitté..")
			os.Exit(0)
		default:
			fmt.Println("Le choix est invalide, réessayer !")
		}
	}
}
