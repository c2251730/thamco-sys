package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"math/rand"
)

type Order struct {
	ID           string    `json:"id"`
	CustomerID   string    `json:"customerId"`
	Products     []Product `json:"products"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeliveryInfo DeliveryInfo `json:"deliveryInfo"`
}

type Product struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type DeliveryInfo struct {
	Address       string `json:"address"`
	PhoneNumber   string `json:"phoneNumber"`
}

var orders = make(map[string]Order)

func main() {
	http.HandleFunc("/orders", createOrderHandler)
	http.HandleFunc("/orders/{orderId}", getOrderHandler)
	http.HandleFunc("/orders/{orderId}/cancel", cancelOrderHandler)

	fmt.Println("Order Microservice is running on port 8082...")
	http.ListenAndServe(":8082", nil)
}

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var newOrder Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	orderID := fmt.Sprintf("%d", time.Now().UnixNano())
	newOrder.ID = orderID
	newOrder.CreatedAt = time.Now()
	newOrder.UpdatedAt = time.Now()

	if checkProductAvailability(newOrder.Products) {
		orders[orderID] = newOrder

		updateStock(newOrder.Products)

		sendEmailNotification(newOrder)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newOrder)
	} else {
		http.Error(w, "One or more products are out of stock", http.StatusBadRequest)
	}
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	orderID := r.URL.Query().Get("orderId")
	order, ok := orders[orderID]
	if !ok {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func cancelOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	orderID := r.URL.Query().Get("orderId")
	order, ok := orders[orderID]
	if !ok {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	order.Status = "Canceled"
	order.UpdatedAt = time.Now()
	orders[orderID] = order

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func someFunction() {
    for _, product := range products {
        fmt.Println(product)
    }
}

func updateStock(products []Product) {
}

func sendEmailNotification(order Order) {
	fmt.Printf("Email notification sent for order %s\n", order.ID)
}
