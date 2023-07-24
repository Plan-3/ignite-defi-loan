package main

import (
    "context"
    "fmt"
    "log"

    // Importing the general purpose Cosmos blockchain client
    "github.com/ignite/cli/ignite/pkg/cosmosclient"

    // Importing the types package of your blog blockchain
    "loan/x/loan/types"
)

func main() {
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

    // Define a message to create a post
    msg := &types.MsgRequestLoan{
            Creator: addr,
            Amount: "100zusd",
			Fee: "1zusd",
			Collateral: "100ctz",
			Deadline: "500",
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