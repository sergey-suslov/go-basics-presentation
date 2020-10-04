package main

import (
	"fmt"
)

type Human struct {
	name string
}

func main() {
	fmt.Println("--------------------------")

	var humanById map[string]Human // ~ nil by default
	fmt.Println(humanById)
	fmt.Println(humanById == nil)
	//humanById["1"] = Human{"Jeff"}

	fmt.Println("--------------------------")

	humanById = make(map[string]Human)
	humanById["1"] = Human{"Jeff"}
	humanById["2"] = Human{"Eric"}
	fmt.Println(humanById)

	fmt.Println("--------------------------")

	shortDeclarationMap := map[string]Human{
		"10": {"Dave"},
		"11": {"Paul"},
	}
	fmt.Println(shortDeclarationMap)

	fmt.Println("--------------------------")

	dave, ok := shortDeclarationMap["11"]
	fmt.Println("Dave =", dave, "Got Dave =", ok)

	fmt.Println("--------------------------")

	// No ordering
	shortDeclarationMap["9"] = Human{"Bob"}
	for i, v := range shortDeclarationMap{
		fmt.Printf("[%s]=%v\n", i, v)
	}
}
