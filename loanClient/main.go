package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/Plan-3/ignite-defi-loan/pkg/routes"

)

func main() {
    r := mux.NewRouter()
	r.HandleFunc("/requestloan", routes.CreateLoan).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}