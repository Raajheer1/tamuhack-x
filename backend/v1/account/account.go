package account

import "net/http"

type Response struct {
	AAvantage string `json:"aadvantage_number"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {

}
