package utils 

import (
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"net/http"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"os"
	

	"loan/x/loan/types"
)

var apiKeys = map[string]interface{}{"cosmos1yq0tjx2d5zj2x4": "key1", "cosmos1yq0tjx2d5zj3x6": "key2"}
type Key struct {
	Key string `json:"key"`
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func FilterLoanByStateApprove(loan *types.QueryAllLoanResponse) []types.Loan {
	var filteredLoan []types.Loan

	for _, v := range loan.Loan {
		if v.State == "requested" {
			
			filteredLoan = append(filteredLoan, v)
		}
	}

	return filteredLoan
}

func FilterLoanByStateLiquidate(loan *types.QueryAllLoanResponse) []types.Loan {
	var filteredLoan []types.Loan

	for _, v := range loan.Loan {
		if v.State == "approved" {
			
			filteredLoan = append(filteredLoan, v)
		}
	}

	return filteredLoan
}

func FilterLoanById(loan *types.QueryAllLoanResponse, account string) []types.Loan {
	var filteredLoan []types.Loan

	for _, v := range loan.Loan {
		if v.Borrower == account{
			filteredLoan = append(filteredLoan, v)
		}
	}

	return filteredLoan
}

func FilterLoanByIdAndState(loan *types.QueryAllLoanResponse, account string) []types.Loan {
	var filteredLoan []types.Loan

	for _, v := range loan.Loan {
		if v.Borrower == account && v.State == "approved"{
			filteredLoan = append(filteredLoan, v)
		}
	}

	return filteredLoan
}

func ApiMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Authorization")
		if apiKey == "" {
				http.Error(w, "API key is missing", http.StatusUnauthorized)
				return
		}

		if _, valid := apiKeys[apiKey]; !valid {
				http.Error(w, "Invalid API key", http.StatusUnauthorized)
				return
		}

		// If the API key is valid, you can also extract any associated data and pass it to the request context if needed.

		// Call the next handler in the chain.
		next.ServeHTTP(w, r)
})
}

func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey
	privKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	// returns []byte need to convert to string
	// privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	// publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)
	// parsedPrivate, _ := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	// parsedPublic, _  := x509.ParsePKCS1PublicKey(publicKeyBytes)

	if err := pem.Encode(os.Stdout, privKeyBlock); err != nil {
		fmt.Println(err)
	}
	
	// fmt.Println("Private Key: ", privKeyBlock)
	// fmt.Println("Pub Key: ", parsedPublic)
	return privateKey, publicKey, nil
	}
