package main

import (
	"fmt"
)

type Named interface {
	PrintName()
}

type Human struct {
	Name string
	Age  uint
}

func (h Human) PrintName() {
	fmt.Println("Print human's name", h.Name)
}

func (h Human) PrintAge() {
	fmt.Println("Print human's age", h.Age)
}

type Developer struct {
	Login string
	Human
	Manager *Human
	Named   Named
}

func (d Developer) PrintAge() {
	fmt.Println("Print developer's age", d.Age)
}

func PrintHuman(h Human) {
	fmt.Println("Human", h)
}

func main() {
	fmt.Println("--------------------------")

	jeff := Human{
		Name: "Jeff",
		Age:  10,
	}

	jeff.PrintName()
	jeff.PrintAge()

	fmt.Println("--------------------------")

	developerPaul := Developer{
		Login: "paul.brown",
		Human: Human{
			Name: "Paul",
			Age:  20,
		},
		Manager: &jeff,
		Named:   jeff,
	}

	developerPaul.PrintName()
	developerPaul.PrintAge()
	developerPaul.Human.PrintAge()
	developerPaul.Named.PrintName()
	developerPaul.Manager.PrintAge()
	fmt.Println("Developer's name direct", developerPaul.Name)

	fmt.Println("--------------------------")

	PrintHuman(developerPaul.Human)
	//PrintHuman(developerPaul)
}
