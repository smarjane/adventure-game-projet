package adventure

import "fmt"

func SpellBook(c1 *Character) {
	// Vérifier si "Boule de feu" existe déjà
	for _, s := range c1.Skills {
		if s == "Boule de feu" {
			fmt.Println("Vous connaissez déjà le sort Boule de feu")
			return
		}
	}
	// Ajouter le sort sans afficher de message
	c1.Skills = append(c1.Skills, "Boule de feu")
}
