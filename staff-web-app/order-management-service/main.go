package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Order struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customerID"`
	Products   []Product `json:"products"`
	Total      float64   `json:"total"`
	Status     string    `json:"status"`
	Dispatched bool      `json:"dispatched"`
	DispatchedAt time.Time `json:"dispatchedAt"`
}

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type OrderManager struct {
	orders map[string]Order
	mu     sync.RWMutex
}

func NewOrderManager() *OrderManager {
	return &OrderManager{
		orders: make(map[string]Order),
	}
}

func (om *OrderManager) PlaceOrder(order Order) error {
	om.mu.Lock()
	defer om.mu.Unlock()

	order.ID = generateUniqueID()

	order.Status = "Processing"
	order.Dispatched = false

	om.orders[order.ID] = order

	return nil
}

func (om *OrderManager) DispatchOrder(orderID string) error {
	om.mu.Lock()
	defer om.mu.Unlock()

	order, found := om.orders[orderID]
	if !found {
		return fmt.Errorf("order not found")
	}

	order.Dispatched = true
	order.DispatchedAt = time.Now()
	order.Status = "Dispatched"

	om.orders[orderID] = order

	return nil
}

func (om *OrderManager) GetOrders() []Order {
	om.mu.RLock()
	defer om.mu.RUnlock()

	var orders []Order
	for _, order := range om.orders {
		orders = append(orders, order)
	}

	return orders
}

func generateUniqueID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func HandlePlaceOrder(w http.ResponseWriter, r *http.Request, om *OrderManager) {
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := om.PlaceOrder(order)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error placing order: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Order placed successfully")
}

func HandleDispatchOrder(w http.ResponseWriter, r *http.Request, om *OrderManager) {
	orderID := r.URL.Query().Get("orderID")
	if orderID == "" {
		http.Error(w, "OrderID is required", http.StatusBadRequest)
		return
	}

	err := om.DispatchOrder(orderID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error dispatching order: %s", err), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Order dispatched successfully")
}

func HandleGetOrders(w http.ResponseWriter, r *http.Request, om *OrderManager) {
	orders := om.GetOrders()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func main() {
	orderManager := NewOrderManager()

	http.HandleFunc("/orders/place", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			HandlePlaceOrder(w, r, orderManager)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/orders/dispatch", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			HandleDispatchOrder(w, r, orderManager)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			HandleGetOrders(w, r, orderManager)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
