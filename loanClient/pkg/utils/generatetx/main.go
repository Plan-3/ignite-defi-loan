package main 

import (
	"context"
    "fmt"

    "google.golang.org/grpc"

    "github.com/cosmos/cosmos-sdk/types/tx"
	"cosmossdk.io/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
)

func sendTx(ctx context.Context) error {
	// Choose your codec: Amino or Protobuf. Here, we use Protobuf, given by the following function.
	app := simapp.NewSimApp(Protobuf)

	// Create a new TxBuilder.
	txBuilder := app.TxConfig().NewTxBuilder()

	// --snip--

	priv1, _, addr1 := testdata.KeyTestPubAddr()
	priv2, _, addr2 := testdata.KeyTestPubAddr()
	priv3, _, addr3 := testdata.KeyTestPubAddr()

	msg1 := banktypes.NewMsgSend(addr1, addr3, types.NewCoins(types.NewInt64Coin("atom", 12)))
	msg2 := banktypes.NewMsgSend(addr2, addr3, types.NewCoins(types.NewInt64Coin("atom", 34)))

	err := txBuilder.SetMsgs(msg1, msg2)
	if err != nil {
			return err
	}

	txBuilder.SetGasLimit(2000000)
	txBuilder.SetFeeAmount(sdk.ParseCoinsNormalized("200usdc"))
	txBuilder.SetMemo("test memo")
	txBuilder.SetTimeoutHeight(999999999)

	privs := []cryptotypes.PrivKey{priv1, priv2}
	accNums:= []uint64{..., ...} // The accounts' account numbers
	accSeqs:= []uint64{..., ...} // The accounts' sequence numbers

	// First round: we gather all the signer infos. We use the "set empty
	// signature" hack to do that.
	var sigsV2 []signing.SignatureV2
	for i, priv := range privs {
			sigV2 := signing.SignatureV2{
					PubKey: priv.PubKey(),
					Data: &signing.SingleSignatureData{
							SignMode:  encCfg.TxConfig.SignModeHandler().DefaultMode(),
							Signature: nil,
					},
					Sequence: accSeqs[i],
			}

			sigsV2 = append(sigsV2, sigV2)
	}
	err := txBuilder.SetSignatures(sigsV2...)
	if err != nil {
			return err
	}

	// Second round: all signer infos are set, so each signer can sign.
	sigsV2 = []signing.SignatureV2{}
	for i, priv := range privs {
			signerData := xauthsigning.SignerData{
					ChainID:       chainID,
					AccountNumber: accNums[i],
					Sequence:      accSeqs[i],
			}
			sigV2, err := tx.SignWithPrivKey(
					encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
					txBuilder, priv, encCfg.TxConfig, accSeqs[i])
			if err != nil {
					return nil, err
			}

			sigsV2 = append(sigsV2, sigV2)
	}
	err = txBuilder.SetSignatures(sigsV2...)
	if err != nil {
			return err
	}

	    // Generated Protobuf-encoded bytes.
			txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
			if err != nil {
					return err
			}
	
			// Generate a JSON string.
			txJSONBytes, err := encCfg.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
			if err != nil {
					return err
			}
			txJSON := string(txJSONBytes)

			grpcConn := grpc.Dial(
        "127.0.0.1:9090", // Or your gRPC server address.
        grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
    )
    defer grpcConn.Close()

    // Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
    // service.
    txClient := tx.NewServiceClient(grpcConn)
    // We then call the BroadcastTx method on this client.
    grpcRes, err := txClient.BroadcastTx(
        ctx,
        &tx.BroadcastTxRequest{
            Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
            TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
        },
    )
    if err != nil {
        return err
    }

    fmt.Println(grpcRes.TxResponse.Code) // Should be `0` if the tx is successful

    return nil
}