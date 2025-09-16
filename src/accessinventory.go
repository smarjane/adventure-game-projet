package adventure

import (
	"fmt"
	"strconv"
)

func AccessInventaire(items []Items) {
	for _, item := range items {
		for i := 0; i < item.Quantite; i++ {
			fmt.Println(item.Nom)
		}
	}
}
func RemoveInventory(Inventaire []Items, item string) ([]Items, bool) {
	for i, Items := range Inventaire {
		if Items.Nom == item {
			return append(Inventaire[:1], Inventaire[i+1]), true
		}
	}
	return Inventaire, false
}
func AccessInventory(c1 Character) {
	fmt.Println("----Inventaire----")
	if len(c1.Inventaire) == 0 {
		fmt.Println("L'inventaire est vide")
	} else {
		for _, items := range c1.Inventaire {
			println(items.Nom + " " + strconv.Itoa(items.Quantite))
		}
		var choix int
		fmt.Print("Entrer le numéro de Item pour le choisir ou 0 pour retour")
		fmt.Scanln(&choix)

		if choix > 0 && choix <= len(c1.Inventaire) {
			selected := c1.Inventaire[choix-1]
			if selected.Nom == "Potion" {
				TakePot(c1)
			}
		} else {
			fmt.Println("Cet élément n'est pas encore pris en charge. ")
		}
	}
}
