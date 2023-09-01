package routes

import (
	"context"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	//"github.com/Plan-3/ignite-defi-loan/loanClient/pkg/utils"

	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/ignite/pkg/cosmosclient"

	// Importing the types package of your blog blockchain
	"loan/x/loan/types"

)

func Stake(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
	}
	defer r.Body.Close()

	msg := &types.MsgStake{}

	err = json.Unmarshal(body, msg)
	if err != nil {
		http.Error(w, "Failed to unmarshal json", 500)
		return
	}

	ctx := context.Background()
	addressPrefix := "cosmos"

	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		log.Print(err)
	}

	account, err := client.Account(msg.Creator)
	if err != nil {
		log.Print(err)
	}

	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		log.Print(err)
	}

	_ = txResp

	w.WriteHeader(http.StatusOK)
}