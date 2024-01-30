package handlers

import (
	"accounts/data"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Account Handler interface
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

// Add a new account
func (ah *AccountHandler) AddAccount(rw http.ResponseWriter, r *http.Request) {
	newAccount := &data.Account{}
	err := newAccount.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to read account from json", http.StatusBadRequest)
	}
	data.AddAccount(newAccount)
}

// GetAccount by ID
func (ah *AccountHandler) GetAccountWithId(rw http.ResponseWriter, r *http.Request) {
	accountId, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(rw, "Account Id is required", http.StatusBadRequest)
	}
	account := data.GetAccount(accountId)

	if account == nil {
		http.Error(rw, "Account with id "+accountId+" not found", http.StatusNotFound)
	}
	rw.WriteHeader(http.StatusOK)
	err := account.ToJson(rw)
	if err != nil {
		ah.l.Println("Error coverting account to json")
		http.Error(rw, "Error converting to json", http.StatusInternalServerError)
	}
}

// Update Account
func (ah *AccountHandler) UpdateAccount(rw http.ResponseWriter, r *http.Request) {
	accountId, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(rw, "Account Id is required", http.StatusBadRequest)
	}

	acct := &data.Account{}
	acct.FromJson(r.Body)
	acct, err := data.UpdateAccount(accountId, acct)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	rw.WriteHeader(http.StatusAccepted)
	acct.ToJson(rw)

}
