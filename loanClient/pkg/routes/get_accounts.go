package routes

import (
	"context"
	"log"
	"net/http"
	"encoding/json"
	//"strconv"

	"github.com/gorilla/mux"
	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/cosmos/cosmos-sdk/types/query"

)

func GetAccountBalances(w http.ResponseWriter, r *http.Request) {

	// use mux to grab var address from url *name is set in main.go
	vars := mux.Vars(r)
	address := vars["address"]


	// Set up your cosmos client and other initialization code here as before...
	ctx := context.Background()
	addressPrefix := "cosmos"

	
	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		log.Fatal(err)
	}

	// instantiate a pointer to query.pagerequest default is 1 default limit is 100
	page := &query.PageRequest{}


	balances, err := client.BankBalances(ctx, address, page)
	if err != nil {
		log.Fatal(err)
	}

	res, _ := json.Marshal(balances)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}