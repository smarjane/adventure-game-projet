package adventure

import (
	"fmt"
)

func DisplayInfo(c1 *Character) {
	fmt.Println("Nom :", c1.Nom)
	fmt.Println("Classe :", c1.Classe)
	fmt.Println("Niveau :", c1.Niveau)
	fmt.Println("Points de vie maximum :", c1.PVMax)
	fmt.Println("Points de vie actuels :", c1.PVActuels)
	fmt.Println("Items de l'inventaire :", c1.Inventaire)
	fmt.Println("Skills:", c1.Skills)
}
