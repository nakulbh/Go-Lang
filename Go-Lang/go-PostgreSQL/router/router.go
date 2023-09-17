package main

import(
	"fmt"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET","OPTIONS")
	router.HandleFunc("/")
	router.HandleFunc("/")
	router.HandleFunc("/")
	router.HandleFunc("/")

}

