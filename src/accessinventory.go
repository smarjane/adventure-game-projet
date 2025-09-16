package adventure

import (
	"fmt"
)

func AccessInventory(items []Items) {
	for _, item := range items {
		for i := 0; i < item.Quantite; i++ {
			fmt.Println(item.Nom)
		}
	}
}
