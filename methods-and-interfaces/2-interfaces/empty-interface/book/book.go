package book

type Book struct {
	Name   string
	Price  float32
	Author string
}

func (b *Book) GetName() string {
	if b == nil {
		return "Nil name"
	}
	return b.Name
}

func (b *Book) GetPrice() float32 {
	return b.Price
}

func (b *Book) GetProducer() string {
	return b.Author
}

//func (b Book) GetName() string {
//	return b.Name
//}
//
//func (b Book) GetPrice() float32 {
//	return b.Price
//}
//
//func (b Book) GetProducer() string {
//	return b.Author
//}

type Product interface {
	GetName() string
	GetPrice() float32
	GetProducer() string
}
