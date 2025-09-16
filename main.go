package main

import (
	adventure "adventure/src"
)

func main() {
	var Items []adventure.Items
	item1 := adventure.InitItems("potion", 3)
	item2 := adventure.InitItems("épée glacial", 2)
	Items = append(Items, item1, item2)
	c1 := adventure.InitCharacter("baltazar", "guerrier", 4, 100, 15, Items)
	adventure.DisplayInfo(c1)
	adventure.AccessInventory(Items)
}
