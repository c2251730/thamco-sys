package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Customer struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Funds    float64 `json:"funds"`
	Orders   []Order `json:"orders"`
}

type Order struct {
	ID      string      `json:"id"`
	Products []Product   `json:"products"`
	Total   float64     `json:"total"`
}

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type CustomerManager struct {
	customers map[string]Customer
	mu        sync.RWMutex
}

func NewCustomerManager() *CustomerManager {
	return &CustomerManager{
		customers: make(map[string]Customer),
	}
}

func (cm *CustomerManager) GetCustomer(customerID string) (Customer, error) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	customer, found := cm.customers[customerID]
	if !found {
		return Customer{}, fmt.Errorf("customer not found")
	}
	return customer, nil
}

func (cm *CustomerManager) DeleteCustomer(customerID string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	customer, found := cm.customers[customerID]
	if !found {
		return fmt.Errorf("customer not found")
	}

	customer.Email = ""


	delete(cm.customers, customerID)

	return nil
}

func HandleGetCustomer(w http.ResponseWriter, r *http.Request, cm *CustomerManager) {
	customerID := r.URL.Query().Get("customerID")
	if customerID == "" {
		http.Error(w, "CustomerID is required", http.StatusBadRequest)
		return
	}

	customer, err := cm.GetCustomer(customerID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func HandleDeleteCustomer(w http.ResponseWriter, r *http.Request, cm *CustomerManager) {
	customerID := r.URL.Query().Get("customerID")
	if customerID == "" {
		http.Error(w, "CustomerID is required", http.StatusBadRequest)
		return
	}

	err := cm.DeleteCustomer(customerID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Customer account deleted successfully")
}

func main() {
	customerManager := NewCustomerManager()

	http.HandleFunc("/customer", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			HandleGetCustomer(w, r, customerManager)
		case http.MethodDelete:
			HandleDeleteCustomer(w, r, customerManager)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
