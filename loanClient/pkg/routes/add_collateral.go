package routes

import (
	"context"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	//"github.com/Plan-3/ignite-defi-loan/loanClient/pkg/utils"

	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/ignite/pkg/cosmosclient"

	// Importing the types package of your blog blockchain
	"loan/x/loan/types"

)

func AddCollateral(w http.ResponseWriter, r *http.Request) {
	// Ensure this endpoint only accepts POST requests
	if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
	}
	defer r.Body.Close()

	// Parse the data from the request body into a loan request struct
	msg := &types.MsgAddCollateral{
		/*
		Creator: loanReq.Creator,
		Id: loanReq.Id,
		Amount: loanReq.Amount
		*/
}

	// Unmarshal the JSON data into the protobuf-generated struct
	// Unmarshal takes byte[] and a proto message
	// turns the byte[] into a proto message
	// do not need all the fields in the body of request to be present
	// does need the json key to match exactly the proto field name
	err = json.Unmarshal(body, msg)
	if err != nil {
		http.Error(w, "Failed to unmarshal json", 500)
		return
	}	
	

	// Set up your cosmos client and other initialization code here as before...
	ctx := context.Background()
	addressPrefix := "cosmos"

	
	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		fmt.Print(err)
	}
	
	
	// Get account from the keyring
	account, err := client.Account(msg.Creator)
	if err != nil {
		fmt.Print(err)
	}

	// Broadcast a transaction from account `msg.creator` with the message
  // to create a post store response in txResp
  txResp, err := client.BroadcastTx(ctx, account, msg)
  if err != nil {
      fmt.Print(err)
			http.Error(w, "Tx failed", 500)
    }
	
	res, _ := json.Marshal(txResp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

 
}