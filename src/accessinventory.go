package adventure

import (
	"fmt"
	"strconv"
	"strings"
)

func AccessInventaire(c1 *Character) {
	fmt.Println("---Inventaire---")
	if len(c1.Inventaire) == 0 {
		fmt.Println("L'inventaire est vide")
	} else {
		for _, items := range c1.Inventaire {
			println(items.Nom + " " + strconv.Itoa(items.Quantite))
		}
		var choix int
		fmt.Print("Entrer le numéro de Item pour le choisir ou 0 pour retour :")
		fmt.Scanln(&choix)

		if choix > 0 && choix <= len(c1.Inventaire) {
			selected := c1.Inventaire[choix-1]
			if strings.ToLower(selected.Nom) == "potion" {
				TakePot(c1)
			}
		} else if choix != 0 {
			fmt.Println("Cet élément n'est pas encore pris en charge.")
		}
	}
}
func RemoveInventaire(Inventaire []Items, item string) ([]Items, bool) {
	for i, Items := range Inventaire {
		if Items.Nom == item {
			fmt.Printf("%s a été ajouté dans l'inventaire. Quantité actuel: %d. \n", Items.Nom)
			return append(Inventaire[:1], Inventaire[i+1]), true
		}
	}
	return Inventaire, false
}
func AddInventaire(c1 *Character, item Items) {
	for i, it := range c1.Inventaire {
		if it.Nom == item.Nom {
			c1.Inventaire[i].Quantite += item.Quantite
			fmt.Printf("%s a été ajouté dans l'inventaire. Quantité actuel: %d.\n", item.Nom, c1.Inventaire[i].Quantite)
			return
		}
	}
	c1.Inventaire = append(c1.Inventaire, item)
	fmt.Printf("%s  dans l'inventaire", item.Nom)
}
