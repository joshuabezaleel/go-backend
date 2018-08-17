package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Product structure description
type Product struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Price    string `json:"price,omitempty"`
	Quantity string `json:"quantity,omitempty"`
}

var products []Product

func main() {
	router := mux.NewRouter()

	products = append(products, Product{ID: "1", Name: "monitor", Price: "1,958,291", Quantity: "21"})
	products = append(products, Product{ID: "2", Name: "mouse", Price: "232,113", Quantity: "32"})
	products = append(products, Product{ID: "3", Name: "laptop", Price: "5,372,123", Quantity: "15"})

	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/product/{id}", CreateProduct).Methods("POST")
	router.HandleFunc("/product/{id}", DeleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// GetProducts function description
func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

// GetProduct function description
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// CreateProduct function description
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = params["id"]
	products = append(products, product)
	json.NewEncoder(w).Encode(products)
}

// DeleteProduct function description
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}
