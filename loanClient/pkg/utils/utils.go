package utils 

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	

	"loan/x/loan/types"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func FilterLoanByState(loan *types.QueryAllLoanResponse) []types.Loan {
	var filteredLoan []types.Loan

	for _, v := range loan.Loan {
		if v.State == "requested" {
			
			filteredLoan = append(filteredLoan, v)
		}
	}

	return filteredLoan
}