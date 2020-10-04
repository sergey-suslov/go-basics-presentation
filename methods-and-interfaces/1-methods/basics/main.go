package main

import (
	"fmt"
	"go-basics-presentation/methods-and-interfaces/1-methods/basics/book"
)

// Can not define methods of non-local types
//func (b book.Book) printInfo() {
//	fmt.Printf("Info %+v\n", b)
//}

type Price float32

func (p Price) printPrice() {
	fmt.Printf("The price is %f\n", p)
}

func (p *Price) setPrice(price Price) {
	*p = price
}

func main() {
	fmt.Println("--------------------------")

	jeff := book.Human{Name: "Jeff"}
	jeff.Greet("Dave")

	fmt.Println("--------------------------")

	price := Price(1)
	price.printPrice()
	//(&price).setPrice(Price(10))
	price.setPrice(Price(10))
	price.printPrice()

	fmt.Println("--------------------------")

	history := book.NewBook("History")
	fmt.Printf("%+v\n", history)

	fmt.Println("--------------------------")

	history.Price = 100
	fmt.Printf("%+v\n", history)

	fmt.Println("--------------------------")

	math := book.NewBook(
		"Math",
		book.SetAuthor(&book.Human{Name: "John Brown"}),
		book.SetMeta("Meta info"),
	)
	fmt.Printf("%+v\n", math)
}
