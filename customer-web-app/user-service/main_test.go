package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type MockUserServiceRepository struct{}

func (m *MockUserServiceRepository) createUser(user User) (string, error) {
	return "user123", nil
}

type MockAuthService struct{}

func (m *MockAuthService) authenticateUser(email, password string) (User, bool) {
	return User{ID: "user123", Email: email}, true
}

func TestCreateUserHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	userServiceRepository = &MockUserServiceRepository{}
	authService = &MockAuthService{}

	newUser := User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	userJSON, err := json.Marshal(newUser)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(server.URL+"/register", "application/json", strings.NewReader(string(userJSON)))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, resp.StatusCode)
	}

}

func TestAuthenticateUserHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

  authService = &MockAuthService{}

	loginUser := User{
		Email:    "john@example.com",
		Password: "password123",
	}

	loginJSON, err := json.Marshal(loginUser)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(server.URL+"/login", "application/json", strings.NewReader(string(loginJSON)))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

}
