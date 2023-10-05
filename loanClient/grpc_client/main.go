package main

import (
"log"
//"github.com/ignite/cli/ignite/pkg/cosmosclient"
"context"
"google.golang.org/grpc"
"google.golang.org/grpc/credentials/insecure"
pb "loan/x/loan/types"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
			log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMsgClient(conn)
	// Set up your cosmos client and other initialization code here as before...
ctx := context.Background()
//addressPrefix := "cosmos"


// // Create a Cosmos client instance
// client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
// if err != nil {
// 	log.Fatal(err)
// }
req := &pb.MsgRequestLoan{
	Creator: "cosmos12mzup4c97w833hy37syxkvk2g4zxy20ky5kqc5",
	Amount: "100zusd",
	Fee: "75usdc",
	Collateral: "12cqt",
	Deadline: "100000",
	}

	r, err := c.RequestLoan(ctx, req)
	if err != nil {
			log.Fatalf("could not greet: %v", err)
	}
	// account, _ := client.Account("cosmos12mzup4c97w833hy37syxkvk2g4zxy20ky5kqc5")

	// txResp, err := client.BroadcastTx(ctx, account, r)
	// if err != nil {
	// 		log.Fatal(err)
	// }
	// log.Printf("Client txresp: %s", txResp)
	log.Printf("C response: %s", r)
}