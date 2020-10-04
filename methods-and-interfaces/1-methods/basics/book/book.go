package book

import "fmt"

type Human struct {
	Name string
}

func (h Human) Greet(name string) {
	fmt.Println("Hi,", name)
}

type Book struct {
	Name   string
	Price  float32
	author *Human
	meta   string
}

type Option func(book *Book)

func NewBook(name string, opts ...Option) *Book {
	book := &Book{Name: name}

	for _, opt := range opts {
		opt(book)
	}

	return book
}

func SetAuthor(author *Human) Option {
	return func(book *Book) {
		book.author = author
	}
}

func SetMeta(meta string) Option {
	return func(book *Book) {
		book.meta = meta
	}
}
