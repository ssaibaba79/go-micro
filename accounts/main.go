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
	accountsRouter.HandleFunc("", handlers.NewAccountsHandler(l).GetAccounts).
		Methods(http.MethodGet)
}
