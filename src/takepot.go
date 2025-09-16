package adventure

import "fmt"

func TakePot(c1 Character) {
	// Vérifier la potion dans l'inventaire
	potion := -1
	for _, Items := range c1.Inventaire {
		if Items.Nom == "potion" {
			potion = 0
			break
		}
		if potion != -1 {
			// Guérir HP de Charactc1r
			c1.PVActuels += 50
			// L'assurance HP Currence dépasse pas Max HP
			if c1.PVActuels >= c1.PVMax {
				c1.PVActuels = c1.PVMax
			}
			// Supprimer potion a été utilisé dans l'inventaire
			c1.Inventaire = append(c1.Inventaire[:potion], c1.Inventaire[potion+1:]...)
			fmt.Println("Vous avez pris un potion et ajouté 50HP")
		} else {
			fmt.Println("Vous avez aucune potion")
		}
	}

}
