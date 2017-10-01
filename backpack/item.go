package backpack

type Item struct {
	name   string
	weight float64
	price  float64
}

func NewItem(name string, weight, price float64) *Item {
	return &Item{name, weight, price}
}

func (item *Item) GetName() string {
	return item.name
}

func (item *Item) GetWeight() float64 {
	return item.weight
}

func (item *Item) GetPrice() float64 {
	return item.price
}
