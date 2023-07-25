package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/Plan-3/ignite-defi-loan/loanClient/pkg/utils"

	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/ignite/pkg/cosmosclient"

	// Importing the types package of your blog blockchain
	"loan/x/loan/types"

)

func CreateLoan(w http.ResponseWriter, r *http.Request) {
	// Ensure this endpoint only accepts POST requests
	if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
	}

	// Parse the data from the request body into a loan request struct
	var loanReq types.MsgRequestLoan
	if err := json.NewDecoder(r.Body).Decode(&loanReq); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
	}

	fmt.Println(loanReq)

	// Set up your cosmos client and other initialization code here as before...
	ctx := context.Background()
	addressPrefix := "cosmos"

	
	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		log.Fatal(err)
	}
	
	// Account `alice` was initialized during `ignite chain serve`
	accountName := "alice"
	
	// Get account from the keyring
	account, err := client.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}
	
	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}
	
	// Your loan creation code here...
	msg := &types.MsgRequestLoan{
		Creator: loanReq.Creator,
		Amount: loanReq.Amount,
		Fee: loanReq.Fee,
		Collateral: loanReq.Collateral,
		Deadline: loanReq.Deadline,
}
	// Broadcast a transaction from account `alice` with the message
    // to create a post store response in txResp
    txResp, err := client.BroadcastTx(ctx, account, msg)
    if err != nil {
        log.Fatal(err)
    }

    // Print response from broadcasting a transaction
    fmt.Print("MsgCreateLoan:\n\n")
    fmt.Println(txResp)

    // Instantiate a query client for your `blog` blockchain
    queryClient := types.NewQueryClient(client.Context())

    // Query the blockchain using the client's `PostAll` method
    // to get all posts store all posts in queryResp
    queryResp, err := queryClient.LoanAll(ctx, &types.QueryAllLoanRequest{})
    if err != nil {
        log.Fatal(err)
    }

    // Print response from querying all the posts
    fmt.Print("\n\nAll loans:\n\n")
    fmt.Println(queryResp)
}