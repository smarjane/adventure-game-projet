package adventure

import "fmt"

func ItemsMax(c1 *Character, nouvelItem Items) bool {
	total := 0
	for _, item := range c1.Inventaire {
		total += item.Quantite
	}

	total += nouvelItem.Quantite

	if total > 10 {
		fmt.Println("Inventaire plein ! Impossible d'ajouter l'objet.")
		return false // Plus de 10 objets, on refuse
	}
	fmt.Println("Ajout accepté :", nouvelItem.Nom)
	return true // C’est bon, on peut ajouter

}
