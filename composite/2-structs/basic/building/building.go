package building

type House struct {
	Id int
	Name string
	Area int
	meta string
}

func NewHouse(id int, name string, area int, meta string) House {
	return House{Id: id, Name: name, Area: area, meta: meta}
}

