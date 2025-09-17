package adventure

import "fmt"

func TakePot(c1 *Character) {
	for i, item := range c1.Inventaire {
		if item.Nom == "potion" && item.Quantite > 0 {
			// Guérir HP du Character
			c1.PVActuels += 50
			if c1.PVActuels > c1.PVMax {
				c1.PVActuels = c1.PVMax
			}

			// Retirer une potion
			c1.Inventaire[i].Quantite--

			fmt.Println("Vous avez pris une potion et ajouté 50 HP")
			return
		}
	}

	fmt.Println("Vous n'avez aucune potion")
}
