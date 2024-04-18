package main

import (
	"encoding/json"
	"net/http"
)

var (
	productStore = make(map[string]Product)
)

func getProduct(w http.ResponseWriter, r *http.Request) {

	productName := r.URL.Query().Get("name")
	if product, ok := productStore[productName]; ok {
		json.NewEncoder(w).Encode(product)
	} else {
		http.Error(w, "Product not found", http.StatusNotFound)
	}
}

func createOrUpdateProduct(w http.ResponseWriter, r *http.Request) {

	var prod Product
	if err := json.NewDecoder(r.Body).Decode(&prod); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productStore[prod.Name] = prod
	json.NewEncoder(w).Encode(prod)
}

func main() {
	http.HandleFunc("/product", getProduct)
	http.HandleFunc("/updateProduct", createOrUpdateProduct)

	http.ListenAndServe(":8080", nil)
}
