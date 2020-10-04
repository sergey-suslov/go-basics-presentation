package main

import (
	"fmt"
	"go-basics-presentation/methods-and-interfaces/2-interfaces/nil-interface/book"
)

func printNilProductName(product book.Product) {
	fmt.Println("Product =", product)
	fmt.Println("Is product nil =", product == nil)
	fmt.Println("Product name =", product.GetName())
}

func main() {
	fmt.Println("--------------------------")

	var product book.Product
	fmt.Println("Product is", product)
	fmt.Println("Is product nil =", product == nil)

	fmt.Println("--------------------------")

	var someBook *book.Book
	product = someBook

	fmt.Println("Book =", someBook)
	printNilProductName(someBook)

	fmt.Println("--------------------------")

	var nilInterface book.Product
	fmt.Printf("Type = %T, value = %v", nilInterface, nilInterface)
}
