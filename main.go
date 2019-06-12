package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/keveaux/go_CRUD_application/apis/product"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/product/findall", product.FindAll).Methods("GET")

	router.HandleFunc("/api/product/search/{keyword}", product.Search).Methods("GET")

	router.HandleFunc("/api/product/searchprices/{min}/{max}", product.SearchPrices).Methods("GET")

	router.HandleFunc("/api/product/create", product.Create).Methods("POST")

	router.HandleFunc("/api/product/update", product.Update).Methods("PUT")

	router.HandleFunc("/api/product/delete/{id}", product.Delete).Methods("DELETE")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
