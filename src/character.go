package adventure

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PVMax      int
	PVActuels  int
	Inventaire []Items
}

type Items struct {
	Nom      string
	Quantite int
}

func InitCharacter(nom string, classe string, niveau int, pvMax int, pvActuels int, Items []Items) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PVMax:      pvMax,
		PVActuels:  pvActuels,
		Inventaire: Items,
	}
}

func InitItems(nom string, quantite int) Items {
	return Items{
		Nom:      nom,
		Quantite: quantite,
	}
}
