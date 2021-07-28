package handlers

import (
	"log"
	"net/http"

	"github.com/PrinceNarteh/microservices-with-go/data"
)

type Product struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Product {
	return &Product{l}
}

//ServeHTTP is the main entry to the handler and satisfies the http.Handler interface
func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handles the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// handles the request for  adding products
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	// catch all
	// if no mothod is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the product list
func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	// fetch the products from the datastore
	productList := data.GetProduct()

	// serialize the list to JSON
	err := productList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(product)
}
