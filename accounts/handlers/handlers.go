package handlers

import (
	"accounts/data"
	"log"
	"net/http"
)

type AccountHandler struct {
	l *log.Logger
}

// Return a AccounHandler
func NewAccountsHandler(l *log.Logger) *AccountHandler {
	return &AccountHandler{l}
}

// GetAccounts
func (ah *AccountHandler) GetAccounts(rw http.ResponseWriter, r *http.Request) {
	ah.l.Println("Handling ", r.Method, r.RequestURI)
	accounts := data.GetAccounts()
	rw.WriteHeader(http.StatusOK)
	err := accounts.ToJson(rw)
	if err != nil {
		ah.l.Println("Error coverting accounts to json")
		http.Error(rw, "Error converting to json", http.StatusInternalServerError)
	}

}
