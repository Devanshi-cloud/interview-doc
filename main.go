package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// User struct
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Product struct
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Home endpoint
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("🚀 Welcome to Go REST API"))
}

// /user endpoint
func getUser(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name: "Devanshi",
		Age:  21,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// /products endpoint (challenge solution)
func getProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{ID: 1, Name: "Laptop", Price: 75000},
		{ID: 2, Name: "Phone", Price: 30000},
		{ID: 3, Name: "Headphones", Price: 2000},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	// Register routes
	http.HandleFunc("/", home)
	http.HandleFunc("/user", getUser)
	http.HandleFunc("/products", getProducts)

	// Start server
	log.Println("✅ Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
