package routes

import (
	"net/http"

	"github.com/Plan-3/ignite-defi-loan/pkg/utils"

)

func GetKey(w http.ResponseWriter, r *http.Request) {
	utils.GenerateRSAKeyPair(2048)
}