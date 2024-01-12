package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var secretKey = []byte("your-secret-key")

type User struct {
	ID       string
	Username string
	Password string
	Role     string
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomClaims struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

var users = []User{
	{ID: "1", Username: "admin", Password: "adminpass", Role: "admin"},
	{ID: "2", Username: "staff", Password: "staffpass", Role: "staff"},
}

func authenticateHandler(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	if err := parseJSONBody(w, r, &credentials); err != nil {
		return
	}

	user, err := authenticateUser(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := createJWTToken(user)
	if err != nil {
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"token": "%s"}`, token)
}

func authenticateUser(username, password string) (*User, error) {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("invalid credentials")
}

func createJWTToken(user *User) (string, error) {
	claims := CustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), 
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func parseJSONBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(v); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return err
	}

	return nil
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/authenticate", authenticateHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
