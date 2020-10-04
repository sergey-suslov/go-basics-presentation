package main

import (
	"fmt"
	"go-basics-presentation/methods-and-interfaces/2-interfaces/nil-interface/book"
)

type Human struct {
	Name string
}

func (h Human) GetName() string {
	return h.Name
}

func printNilProductName(input interface{}) {
	if bookProduct, ok := (input).(book.Product); ok {
		fmt.Println("Product name =", bookProduct.GetName())
	}

	if bookProduct, ok := (input).(book.Book); ok {
		fmt.Println("Book name =", bookProduct.GetName())
	}

	if bookProduct, ok := (input).(*book.Book); ok {
		fmt.Println("*Book name =", bookProduct.GetName())
	}

	if bookProduct, ok := (input).(Human); ok {
		fmt.Println("Human name =", bookProduct.GetName())
	}
	fmt.Printf("----------\n")
}

func main() {
	fmt.Println("--------------------------")

	var emptyInterfaceProduct interface{} = book.Book{Name: "Some book"}
	printNilProductName(emptyInterfaceProduct)
	printNilProductName(&book.Book{Name: "Another book"})
	printNilProductName(Human{"Jeff"})
}
