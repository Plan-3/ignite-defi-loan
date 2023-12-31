package routes

import (
	"context"
	"log"
	"net/http"
	"encoding/json"
	"github.com/Plan-3/ignite-defi-loan/pkg/utils"

	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/ignite/pkg/cosmosclient"

	// Importing the types package of your blog blockchain
	"loan/x/loan/types"

)

func GetLoansRequested(w http.ResponseWriter, r *http.Request) {

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

    // Query the blockchain using the client's `LoanAll` method
    // to get all posts store all posts in queryResp
    queryResp, err := queryClient.LoanAll(ctx, &types.QueryAllLoanRequest{})
    if err != nil {
        log.Fatal(err)
    }

		filtered := utils.FilterLoanByStateApprove(queryResp)

    // Print response from querying all the loans
    // fmt.Print("\n\nAll loans:\n\n")
    // fmt.Printf("%T", queryResp)

		// marshal response back to bytes[] and send to client
		res, _ := json.Marshal(filtered)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
}