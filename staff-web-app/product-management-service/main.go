package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Prices      []Price `json:"prices"`
}

type Price struct {
	Supplier string  `json:"supplier"`
	Price    float64 `json:"price"`
}

type Catalog struct {
	Products []Product `json:"products"`
}

func CalculateFinalPrice(product Product) float64 {
	if len(product.Prices) == 0 {
		return 0.0 
	}

	cheapestPrice := product.Prices[0].Price
	for _, price := range product.Prices {
		if price.Price < cheapestPrice {
			cheapestPrice = price.Price
		}
	}

	finalPrice := cheapestPrice * 1.1
	return finalPrice
}

func updateCatalog() {
	for {
		products := fetchProductsFromSuppliers()

		for i := range products {
			products[i].Prices[0].Price = 100
		}

		time.Sleep(24 * time.Hour)
	}
}

func fetchProductsFromSuppliers() []Product {
	return []Product{
		{
			ID:          "1",
			Name:        "Product 1",
			Description: "Description for Product 1",
			Prices: []Price{
				{Supplier: "SupplierA", Price: 50.0},
				{Supplier: "SupplierB", Price: 45.0},
				{Supplier: "SupplierC", Price: 55.0},
			},
		},
	}
}

func catalogHandler(w http.ResponseWriter, r *http.Request) {
	catalog := fetchProductsFromSuppliers()

	for i := range catalog {
		catalog[i].Prices[0].Price = CalculateFinalPrice(catalog[i])
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(catalog)
}

func main() {
	go updateCatalog()

	http.HandleFunc("/catalog", catalogHandler)
	http.ListenAndServe(":8080", nil)
}
