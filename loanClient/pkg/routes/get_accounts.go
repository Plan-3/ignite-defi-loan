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
	"loan/x/loan/types"
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

	// loop through balances and add price to each coin may need make([]types.TokenPrice, len(balances)) or slice / array
	balancesWithPrice := types.TokenPrice{}
	s := make([]types.TokenPrice, 0)
	for i := 0; i < len(balances); i++ {
		balancesWithPrice.Denom = balances[i]
		switch balances[i].Denom {
		case "ctz":
			balancesWithPrice.Price = 1800;
			break;
		case "cqt":
			balancesWithPrice.Price = 100;
			break;
		default:
			balancesWithPrice.Price = 1;
			break;
		}
		s = append(s, balancesWithPrice)
	}

	res, _ := json.Marshal(s)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}