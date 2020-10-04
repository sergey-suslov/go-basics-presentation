package main

import (
	"fmt"
	"go-basics-presentation/methods-and-interfaces/2-interfaces/basics/book"
)

func productInfoStruct(product book.Product) {
	fmt.Printf("Product interface address = %p\n", &product)
	fmt.Printf("Product type = %T, Product value %+v\n", product, product)

	someBook := (product).(book.Book)

	fmt.Printf("Product casted interface address = %p\n", &someBook)
	someBook.Name = "Edited Name"
}

func productInfoPointer(product book.Product) {
	fmt.Printf("Product interface address = %p\n", &product)
	fmt.Printf("Product type = %T, Product value %+v\n", product, product)

	someBook := (product).(*book.Book)

	fmt.Printf("Product casted interface address = %p\n", someBook)
	someBook.Name = "Edited Name"
}

func productInfoInterface(product *book.Product) {
	fmt.Printf("Product interface address = %p\n", product)
}

func main() {
	fmt.Println("--------------------------")

	someBook := book.Book{Name: "Some book", Price: 100, Author: "Tom"}

	fmt.Printf("Product struct address = %p\n", &someBook)
	productInfoStruct(someBook)
	fmt.Println(someBook)

	fmt.Println("--------------------------")

	fmt.Printf("Product struct address = %p\n", &someBook)
	productInfoPointer(&someBook)
	fmt.Println(someBook)

	fmt.Println("--------------------------")

	var product book.Product = someBook
	fmt.Printf("Product struct address = %p\n", &product)
	productInfoInterface(&product)
	fmt.Println(product)
}
