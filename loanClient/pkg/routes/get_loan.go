package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	//"github.com/Plan-3/ignite-defi-loan/loanClient/pkg/utils"

	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/gorilla/mux"
	// Importing the types package of your blog blockchain
	"loan/x/loan/types"

)

func GetLoan(w http.ResponseWriter, r *http.Request) {
	// use mux.Vars to get the id from the request
	// ["id"] is defined in main 
		queryParams := mux.Vars(r)
		strId := queryParams["id"]

		// convert id from string to int
		id, err := strconv.ParseUint(strId, 10, 64)
		if err != nil {
			http.Error(w, "Failed to convert id to int", 500)
			return
		}

		// Set up your cosmos client and other initialization code here as before...
		ctx := context.Background()
		addressPrefix := "cosmos"
	
		
		// Create a Cosmos client instance
		client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
		if err != nil {
			log.Fatal(err)
		}

    // Instantiate a query client for your `blog` blockchain
    queryClient := types.NewQueryClient(client.Context())

		query := &types.QueryGetLoanRequest{}
		query.Id = id
		fmt.Println(query)
    queryResp, err := queryClient.Loan(ctx, query)
    if err != nil {
        log.Fatal(err)
    }

    // Print response from querying all the posts
    fmt.Print("\n\nAll loans:\n\n")
    fmt.Println(queryResp)
		
		// marshal response back to bytes[] and send to client
		res, _ := json.Marshal(queryResp)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
}