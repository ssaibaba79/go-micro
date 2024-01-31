package main

import (
	"accounts/handlers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// Globals
var (
	l                        = log.New(os.Stdout, "account-api ", log.LstdFlags)
	listenAddress            = "localhost:9090"
	READ_TIMEOUT_IN_SECONDS  = 15 * time.Second
	WRITE_TIMEOUT_IN_SECONDS = 15 * time.Second
)

// Launch Server
func main() {
	router := mux.NewRouter()
	initializeRequestHandlers(router)

	// configure http server
	server := &http.Server{
		Addr:         listenAddress,
		Handler:      router,
		ReadTimeout:  READ_TIMEOUT_IN_SECONDS,
		WriteTimeout: WRITE_TIMEOUT_IN_SECONDS,
		ErrorLog:     l,
	}

	l.Println("Starting server at", listenAddress)
	log.Fatal(server.ListenAndServe())
}

// Initialize request handlers
func initializeRequestHandlers(r *mux.Router) {
	accountsRouter := r.PathPrefix("/accounts").Subrouter()
	accountsHandler := handlers.NewAccountsHandler(l)
	accountsRouter.HandleFunc("", accountsHandler.GetAccounts).
		Methods(http.MethodGet)
	accountsRouter.HandleFunc("", accountsHandler.AddAccount).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/json")
	accountsRouter.HandleFunc("/{id}", accountsHandler.GetAccountWithId).
		Methods(http.MethodGet)
	accountsRouter.HandleFunc("/{id}", accountsHandler.UpdateAccount).
		Methods(http.MethodPut).
		Headers("Content-Type", "application/json")
}
