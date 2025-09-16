package adventure

import (
	"fmt"
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
