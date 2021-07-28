package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

//Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreateOn    string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product

var ErrProductNotFound = fmt.Errorf("Product not found")

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does
// have to buffer the output into an in memory slice of bytes

// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProduct() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

func getNextID() int {
	productListLength := productList[len(productList)-1]
	return productListLength.ID + 1
}

func findProduct(id int) (*Product, int, error) {
	for index, p := range productList {
		if p.ID == id {
			return p, index, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Forthy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreateOn:    time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffer without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreateOn:    time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
