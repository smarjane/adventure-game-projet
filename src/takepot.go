package adventure

import "fmt"

func TakePot(C Character) {
	// Vérifier la potion dans l'inventaire
	potion := -1
	for _, Item := range C.Inventaire {
		if Item.Nom == "potion" {
			potion = 0
			break
		}
		if potion != -1 {
			// Guérir HP de Character
			C.PVActuels += 50
			// L'assurance HP Currence dépasse pas Max HP
			if C.PVActuels >= C.PVMax {
				C.PVActuels = C.PVMax
			}
			// Supprimer potion a été utilisé dans l'inventaire
			C.Inventaire = append(C.Inventaire[:potion], C.Inventaire[potion+1:]...)
			fmt.Println("Vous avez pris un potion et ajouté 50HP")
		} else {
			fmt.Println("Vous avez aucune potion")
		}
	}

}
