package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Id   int
	Name string
}

func customers(w http.ResponseWriter, r *http.Request) {
	c := Customer{Id: 1, Name: "Casper"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c)
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/customers", customers).Methods("GET").Name("customers")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
