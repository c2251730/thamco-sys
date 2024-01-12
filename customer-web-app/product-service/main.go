package main

import (
	"encoding/json"
	"fmt"
        "math/rand"
	"net/http"
)

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

var products = make(map[string]Product)

func main() {
	products["1"] = Product{ID: "1", Name: "Product A", Description: "Description A", Price: 20.99, Stock: 100}
	products["2"] = Product{ID: "2", Name: "Product B", Description: "Description B", Price: 30.49, Stock: 50}

	http.HandleFunc("/products", listProductsHandler)
	http.HandleFunc("/products/updateStock", updateStockHandler)

	fmt.Println("Product Microservice is running on port 8081...")
	http.ListenAndServe(":8081", nil)
}

func listProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func updateStockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	for id := range products {
		newStock := products[id].Stock + randomStockChange()
		if newStock < 0 {
			newStock = 0
		}
		product := products[id]
		product.Stock = newStock
		products[id] = product
	}

	fmt.Println("Stock status updated.")

	w.WriteHeader(http.StatusOK)
}

func randomStockChange() int {
	return -5 + rand.Intn(11)
}


