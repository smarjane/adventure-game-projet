package main

import (
	adventure "adventure/src"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Tâche 1 : Création du personnage
	c1 := adventure.CharacterCreation()

	for {
		fmt.Println("\n--- Menu Principal ---")
		fmt.Println("1: Infos du personnage")
		fmt.Println("2: Inventaire")
		fmt.Println("3: Quitter")

		var choix string
		fmt.Print("Entrez votre choix : ")
		fmt.Scanln(&choix)
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			adventure.ClearTerminal()
			adventure.DisplayInfo(c1)

		case "2":
			adventure.ClearTerminal()
			adventure.AccessInventaire(c1)

			// Sous-menu inventaire
			fmt.Println("\n--- Menu Inventaire ---")
			fmt.Println("1: Utiliser une potion")
			fmt.Println("2: Retour au menu principal")

			var choixInv string
			fmt.Print("Entrez votre choix : ")
			fmt.Scanln(&choixInv)
			choixInv = strings.TrimSpace(choixInv)

			switch choixInv {
			case "1":
				adventure.TakePot(c1)
			case "2":
				// Retour au menu principal
			default:
				fmt.Println("❌ Choix invalide dans le menu inventaire.")
				//adventure.AccessInventaire(c1)
			}

		case "3":
			fmt.Println("👋 Au revoir, aventure terminée.")
			os.Exit(0)

		default:
			fmt.Println("❌ Choix invalide, veuillez réessayer.")
		}
	}
}
