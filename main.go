package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hey Hi, I am a hello handler")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi, I am Mujahid Islam and I am a student of Polytechnic")
}

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

var productsList []Product

func productHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productsList)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", helloHandler)
	mux.HandleFunc("GET /about", aboutHandler)
	mux.HandleFunc("GET /products", productHandler)

	fmt.Println("Server is running on port :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func init() {
	productsList = []Product{
		{Id: 1, Title: "Banana", Description: "Banana is a sweet fruit", Price: 30.5, ImgUrl: "banana.png"},
		{Id: 2, Title: "Apple", Description: "Apple is healthy fruit", Price: 120.0, ImgUrl: "apple.png"},
		{Id: 3, Title: "Orange", Description: "Orange is full of vitamin C", Price: 80.0, ImgUrl: "orange.png"},
		{Id: 4, Title: "Mango", Description: "Mango is king of fruits", Price: 150.0, ImgUrl: "mango.png"},
		{Id: 5, Title: "Pineapple", Description: "Pineapple is juicy fruit", Price: 90.0, ImgUrl: "pineapple.png"},
		{Id: 6, Title: "Grapes", Description: "Grapes are small and sweet", Price: 110.0, ImgUrl: "grapes.png"},
		{Id: 7, Title: "Watermelon", Description: "Watermelon is refreshing", Price: 60.0, ImgUrl: "watermelon.png"},
		{Id: 8, Title: "Papaya", Description: "Papaya helps digestion", Price: 70.0, ImgUrl: "papaya.png"},
		{Id: 9, Title: "Guava", Description: "Guava is rich in fiber", Price: 50.0, ImgUrl: "guava.png"},
		{Id: 10, Title: "Strawberry", Description: "Strawberry is tasty", Price: 200.0, ImgUrl: "strawberry.png"},
	}
}
