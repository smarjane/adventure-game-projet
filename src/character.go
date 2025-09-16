package adventure

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PVMax      int
	PVActuels  int
	Inventaire []Items
	Skills     []string
}

type Items struct {
	Nom      string
	Quantite int
}

func InitCharacter(nom string, classe string, niveau int, pvMax int, pvActuels int, items []Items, skills []string) *Character {
	return &Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PVMax:      pvMax,
		PVActuels:  pvActuels,
		Inventaire: items,
		Skills:     skills,
	}
}

func InitItems(nom string, quantite int) Items {
	return Items{
		Nom:      nom,
		Quantite: quantite,
	}
}
