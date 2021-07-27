package data

import "time"

//Product defines the structure for an API product
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreateOn    string
	UpdatedOn   string
	DeletedOn   string
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Forthy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreateOn:    time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffer without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreateOn:    time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
