package adventure

import (
	"fmt"
	"strconv"
)

func DisplayInfo(c1 Character) {
	fmt.Println("Nom :", c1.Nom)
	fmt.Println("Classe :", c1.Classe)
	fmt.Println("Niveau :", c1.Niveau)
	fmt.Println("Points de vie maximum :", c1.PVMax)
	fmt.Println("Points de vie actuels :", c1.PVActuels)
	fmt.Print("Items de l'inventaire :")
	// Utiliser le boucle pour imprimer chaque Item
	for _, item := range c1.Inventaire {
		//fmt.Printf("(Quantité) %s %d", item.Nom, item.Quantite, "\n")
		print(item.Nom + " " + strconv.Itoa(item.Quantite) + ", ")
	}
}
