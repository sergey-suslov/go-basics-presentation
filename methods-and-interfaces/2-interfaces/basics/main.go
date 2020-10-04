package main

import (
	"fmt"
	"go-basics-presentation/methods-and-interfaces/2-interfaces/basics/book"
)

func printProductInfo(product book.Product) {
	fmt.Println("---------INFO START--------")
	fmt.Println("Product's name", product.GetName())
	fmt.Println("Product's price", product.GetPrice())
	fmt.Println("Product's producer", product.GetProducer())
	fmt.Println("---------INFO END----------")
}

func printProducer(readable book.Readable) {
	fmt.Println("---------PRODUCER START----")
	fmt.Println(readable.GetProducer())
	fmt.Println("---------PRODUCER END------")
}

func main() {
	someBook := book.Book{Name: "Some book", Price: 100, Author: "Tom"}
	someMagazine := book.Magazine{
		Name:      "Some magazine",
		Price:     150,
		Publisher: "The Guardian",
	}

	printProductInfo(someBook)

	printProducer(someBook)
	printProducer(someMagazine)
}
