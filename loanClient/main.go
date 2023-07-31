package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/Plan-3/ignite-defi-loan/pkg/routes"
    "github.com/rs/cors"

)
func handleOptions(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Replace with your Next.js app domain
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.WriteHeader(http.StatusOK)
}

func main() {
    r := mux.NewRouter()
    // TODO: Add query get routes
    r.HandleFunc("/getaccounts/{address}", routes.GetAccountBalances).Methods("GET")
    r.HandleFunc("/getloans", routes.GetLoans).Methods("GET")
    r.HandleFunc("/getloan/{id}", routes.GetLoan).Methods("GET")
	r.HandleFunc("/requestloan", routes.CreateLoan).Methods("POST")
    r.HandleFunc("/cancelloan", routes.CancelLoan).Methods("POST")
    r.HandleFunc("/approveloan", routes.ApproveLoan).Methods("POST")
    r.HandleFunc("/liquidateloan", routes.LiquidateLoan).Methods("POST")
    r.HandleFunc("/repayloan", routes.RepayLoan).Methods("POST")

    // Create a new cors handler with the desired CORS options
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:3000/loan"}, // Replace with your Next.js app domain
		AllowedMethods: []string{"GET", "POST", "OPTIONS"}, // Allow desired HTTP methods
		AllowedHeaders: []string{"*"}, // Allow any headers from the client
	})
    
    r.HandleFunc("/{any:.*}", handleOptions).Methods("OPTIONS")

    handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(":8080", handler))
}