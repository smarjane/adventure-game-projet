package adventure

import "fmt"

func TakePot(c1 *Character) {
	for i, item := range c1.Inventaire {
		if item.Nom == "potion" && item.Quantite > 0 {
			if c1.PVActuels < c1.PVMax {
				c1.PVActuels += 50
				if c1.PVActuels > c1.PVMax {
					c1.PVActuels = c1.PVMax
				}
				c1.Inventaire[i].Quantite--
				fmt.Println("🧪 Vous avez utilisé une Potion !")
				fmt.Printf("❤️ PV actuels : %d\n", c1.PVActuels)
			} else {
				fmt.Println("✅ Vos PV sont déjà au maximum.")
			}
			return
		}
	}
	fmt.Println("❌ Vous n'avez pas de Potion dans votre inventaire.")
}
