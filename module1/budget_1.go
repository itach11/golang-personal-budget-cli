package module1

// Budget stores budget information
type Budget struct {
	Max   float32
	Items []Item
}
type Item struct {
	Description string
	Price       float32
}

// Item stores item information
