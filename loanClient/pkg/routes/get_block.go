package routes

import (
	"context"
	"log"
	"net/http"
	"encoding/json"
	//"strconv"

	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/ignite/pkg/cosmosclient"

)

func GetBlock(w http.ResponseWriter, r *http.Request) {

	// Set up your cosmos client and other initialization code here as before...
	ctx := context.Background()
	addressPrefix := "cosmos"

	
	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.LatestBlockHeight(ctx);
	if err != nil {
		log.Fatal(err)
	}

	res, _ := json.Marshal(block)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}