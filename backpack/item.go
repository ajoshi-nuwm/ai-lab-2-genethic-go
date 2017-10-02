package backpack

// Structure for item - contains name, weight and prices of item
type Item struct {
	name   int
	weight float64
	price  float64
}

// Constructor
func NewItem(name int, weight, price float64) *Item {
	return &Item{name, weight, price}
}

func (item *Item) GetName() int {
	return item.name
}

func (item *Item) GetWeight() float64 {
	return item.weight
}

func (item *Item) GetPrice() float64 {
	return item.price
}

func (item Item) GetValue() float64 {
	return item.GetPrice()
}