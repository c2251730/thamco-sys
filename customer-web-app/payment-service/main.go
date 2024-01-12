package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type PaymentRequest struct {
	OrderID       string  `json:"orderId"`
	CustomerID    string  `json:"customerId"`
	Amount        float64 `json:"amount"`
	CreditCard    string  `json:"creditCard"`
	Expiration    string  `json:"expiration"`
	CVV           string  `json:"cvv"`
}

type PaymentResponse struct {
	OrderID     string  `json:"orderId"`
	Success     bool    `json:"success"`
	Message     string  `json:"message"`
	TransactionID string `json:"transactionId"`
}

func main() {
	http.HandleFunc("/payment/process", processPaymentHandler)

	fmt.Println("Payment Microservice is running on port 8083...")
	http.ListenAndServe(":8083", nil)
}

func processPaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var paymentRequest PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&paymentRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}


	transactionID := generateTransactionID()
	paymentResponse := PaymentResponse{
		OrderID:       paymentRequest.OrderID,
		Success:       true,
		Message:       "Payment successful",
		TransactionID: transactionID,
	}

	updateOrderStatus(paymentRequest.OrderID, "Paid")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(paymentResponse)
}

func updateOrderStatus(orderID, status string) {
	fmt.Printf("Order %s status updated to %s\n", orderID, status)
}

func generateTransactionID() string {
	return fmt.Sprintf("%d", generateRandomNumber(100000, 999999))
}

func generateRandomNumber(min, max int) int {
	return min + generateRandom().Intn(max-min+1)
}

func generateRandom() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
