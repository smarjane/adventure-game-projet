package main

import (
	adventure "adventure/src"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Tâche 1
	Items := []adventure.Items{
		adventure.InitItems("potion", 3),
	}
	c1 := adventure.InitCharacter("baltazar", "guerrier", 4, 100, 45, Items)
	adventure.DisplayInfo(c1)
	adventure.AccessInventory(c1)
	// Tâche6-For : la boucle de menu principal
	for {
		fmt.Println("\n--- Menu Principal ---")
		fmt.Println("1: Info de Charactère")
		fmt.Println("2: L'inventaire")
		fmt.Println("3: Quitter")
		// co the them -- vào sau
		var choix string
		fmt.Println("Entrer votre choix")
		fmt.Scanln(&choix)
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			adventure.DisplayInfo(c1)
		case "2":
			adventure.AccessInventory(c1)
		case "3":
			fmt.Println("En cours quitté..")
			os.Exit(0)
		default:
			fmt.Println("Le choix est invalide, réessayer!!!!")
		}
	}
}
