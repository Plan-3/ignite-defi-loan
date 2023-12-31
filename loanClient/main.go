package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/Plan-3/ignite-defi-loan/pkg/routes"
    //"github.com/Plan-3/ignite-defi-loan/pkg/utils"
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
    r.HandleFunc("/blockheight", routes.GetBlock).Methods("GET")
    r.HandleFunc("/requestkeys", routes.GetKey).Methods("GET")
    r.HandleFunc("/getaccounts/{address}", routes.GetAccountBalances).Methods("GET")
    r.HandleFunc("/getloans", routes.GetLoans).Methods("GET")
    r.HandleFunc("/getloans/requested", routes.GetLoansRequested).Methods("GET")
    r.HandleFunc("/getloans/approved", routes.GetLoansApproved).Methods("GET")
    r.HandleFunc("/getloans/{address}", routes.GetLoansAccount).Methods("GET")
    r.HandleFunc("/getloans/{address}/repay", routes.GetLoansAccountRepay).Methods("GET")
    r.HandleFunc("/getloan/{id}", routes.GetLoan).Methods("GET")
    r.HandleFunc("/redeem", routes.Redeem).Methods("POST")
    r.HandleFunc("/stake", routes.Stake).Methods("POST")
    r.HandleFunc("/withdraw", routes.WithdrawStake).Methods("POST")
	r.HandleFunc("/requestloan", routes.CreateLoan).Methods("POST")
    r.HandleFunc("/cancelloan", routes.CancelLoan).Methods("POST")
    r.HandleFunc("/approveloan", routes.ApproveLoan).Methods("POST")
    r.HandleFunc("/liquidateloan", routes.LiquidateLoan).Methods("POST")
    r.HandleFunc("/repayloan", routes.RepayLoan).Methods("POST")
    r.HandleFunc("/addcollateral", routes.AddCollateral).Methods("POST")
    r.HandleFunc("/withdrawpartial", routes.WithdrawPartial).Methods("POST")
    // check if api key is valid
    // r.Use(utils.ApiMiddleWare)
    // Create a new cors handler with the desired CORS options
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:3000/loan", "http://localhost:3000/account"}, // Replace with your Next.js app domain
		AllowedMethods: []string{"GET", "POST", "OPTIONS"}, // Allow desired HTTP methods
		AllowedHeaders: []string{"*"}, // Allow any headers from the client
	})
    
    r.HandleFunc("/{any:.*}", handleOptions).Methods("OPTIONS")

    handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(":8080", handler))
}