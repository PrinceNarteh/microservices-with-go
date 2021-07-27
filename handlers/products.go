package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PrinceNarteh/microservices-with-go/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (g *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	productList := data.GetProduct()
	data, err := json.Marshal(productList)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	rw.Write(data)
}
