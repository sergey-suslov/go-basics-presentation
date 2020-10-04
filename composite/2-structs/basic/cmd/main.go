package main

import (
	"fmt"
	"go-basics-presentation/composite/2-structs/basic/building"
)

type Human struct {
	name string
	age uint
}

func main() {
	bill := Human{}
	fmt.Println(bill)

	bill.age = 10
	bill.name = "Bill"
	fmt.Println(bill)

	fmt.Println("--------------------------")

	mike := Human{name: "Mike", age: 15}
	fmt.Println(mike)

	fmt.Println("--------------------------")

	jeff := Human{"Jeff",  25}
	fmt.Println(jeff)

	fmt.Println("--------------------------")

	//house := building.House{1, "Green House", 53}
	//fmt.Println(house.meta)
	house := building.NewHouse(1, "Green House", 53, "Meta")
	fmt.Println(house)
}
