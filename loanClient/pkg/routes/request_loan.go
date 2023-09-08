package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"io/ioutil"
	"encoding/json"
	//"github.com/Plan-3/ignite-defi-loan/loanClient/pkg/utils"

	// Importing the general purpose Cosmos blockchain client
	"github.com/ignite/cli/ignite/pkg/cosmosclient"

	// Importing the types package of your blog blockchain
	"loan/x/loan/types"

)

func convertDecimalAmount(amount string) (string, error) {
	// Use regular expression to split the input string into number and string parts
	re := regexp.MustCompile(`^([0-9.]+)([A-Za-z]+)$`)
	matches := re.FindStringSubmatch(amount)

	if len(matches) != 3 {
		return "", fmt.Errorf("Invalid input format")
	}

	decimalPart := matches[1]
	stringPart := matches[2]

	// Parse the decimal part into a float64
	decimalAmount, err := strconv.ParseFloat(decimalPart, 64)
	if err != nil {
		return "", err
	}

	// Multiply the decimal amount by 10^9 and round to an integer
	integerAmount := int64(decimalAmount * 1e9)

	// Convert the integer amount back to a string
	result := fmt.Sprintf("%d%s", integerAmount, stringPart)

	return result, nil
}

func CreateLoan(w http.ResponseWriter, r *http.Request) {
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
	msg := &types.MsgRequestLoan{
		/*
		Creator: loanReq.Creator,
		Amount: loanReq.Amount,
		Fee: loanReq.Fee,
		Collateral: loanReq.Collateral,
		Deadline: loanReq.Deadline, 
		*/
}

  // convert fee to billion units
	// first unmarshal into a json object 
	var request struct {
		Creator string `json:"creator"`
		Amount string `json:"amount"`
		Fee string `json:"fee"`
		Collateral string `json:"collateral"`
		Deadline string `json:"deadline"`
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Failed to unmarshal json into json struct", 500)
	}

	request.Fee, err = convertDecimalAmount(request.Fee)
	if err != nil {
		http.Error(w, "Failed to convert fee", 500)
	}

	body2, err := json.Marshal(request)
	if err != nil {
		http.Error(w, "Failed to marshal json", 500)
	}

	// Unmarshal the JSON data into the protobuf-generated struct
	// Unmarshal takes byte[] and a proto message
	// turns the byte[] into a proto message
	// do not need all the fields in the body of request to be present
	// does need the json key to match exactly the proto field name
	err = json.Unmarshal(body2, msg)
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
		log.Print(err)
	}
	
	// Account `alice` was initialized during `ignite chain serve`
	// can be name or address
	
	/* Get account from the keyring
	accountName := r.Body.Creator
	// Account returns the account with name or address equal to nameOrAddress.
func (c Client) Account(nameOrAddress string) (cosmosaccount.Account, error) {
	defer c.lockBech32Prefix()()

	acc, err := c.AccountRegistry.GetByName(nameOrAddress)
	if err == nil {
		return acc, nil
	}
	return c.AccountRegistry.GetByAddress(nameOrAddress)
}

// Address returns the account address from account name.
func (c Client) Address(accountName string) (string, error) {
	a, err := c.AccountRegistry.GetByName(accountName)
	if err != nil {
		return "", err
	}
	return a.Address(c.addressPrefix)
}


addr, err := account.Address(addressPrefix)
if err != nil {
	log.Print(err)
}
*/
account, err := client.Account(msg.Creator)
if err != nil {
	log.Print(err)
}

	// Broadcast a transaction from account `alice` with the message
  // to create a post store response in txResp
  txResp, err := client.BroadcastTx(ctx, account, msg)
  if err != nil {
      log.Print(err)
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
        log.Print(err)
    }

    // Print response from querying all the posts
    fmt.Print("\n\nAll loans:\n\n")
    fmt.Println(queryResp)
		fmt.Println(ctx)



    // Respond with a 200 status to indicate that preflight request is allowed
    w.WriteHeader(http.StatusOK)
}