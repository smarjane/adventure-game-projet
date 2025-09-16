package adventure

import "fmt"

func TakePot(c1 *Character) {
	// Vérifier la potion dans l'inventaire
	potionIndex := -1
	for i, item := range c1.Inventaire {
		if item.Nom == "potion" {
			potionIndex = i
			break
		}
	}

	if potionIndex != -1 {
		// Guérir HP du Character
		c1.PVActuels += 50
		// S'assurer que PVActuels ne dépasse pas PVMax
		if c1.PVActuels > c1.PVMax {
			c1.PVActuels = c1.PVMax
		}
		// Supprimer la potion utilisée de l'inventaire
		c1.Inventaire = append(c1.Inventaire[:potionIndex], c1.Inventaire[potionIndex+1:]...)
		fmt.Println("Vous avez pris une potion et ajouté 50 HP")
	} else {
		fmt.Println("Vous n'avez aucune potion")
	}
}
