package book

type Magazine struct {
	Name      string
	Price     float32
	Publisher string
}

func (m Magazine) GetProducer() string {
	return m.Publisher
}

type Book struct {
	Name   string
	Price  float32
	Author string
}

func (b Book) GetName() string {
	return b.Name
}

func (b Book) GetPrice() float32 {
	return b.Price
}

func (b Book) GetProducer() string {
	return b.Author
}

type Product interface {
	GetName() string
	GetPrice() float32
	GetProducer() string
}

type Readable interface {
	GetProducer() string
}

// Can't use interface as a receiver
//func (r Readable) PrintProducer() {
//	fmt.Println("Producer is", r.getProducer())
//}
