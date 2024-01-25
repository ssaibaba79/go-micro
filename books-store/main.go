package main

import (
	"book-store/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	bindAddress := ":9090"
	l := log.New(os.Stdout, "book-store", log.LstdFlags)

	r := mux.NewRouter()
	r.Handle("/books", handlers.NewBooks(l))

	s := http.Server{
		Addr:         bindAddress,
		Handler:      r,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// start the server
	go func() {
		l.Println("Starting server on port ", bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	log.Println("Waiting for running tasks to complete...")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	log.Println("Shutting down")

	s.Shutdown(ctx)

}

func StdHttpServer() {

	bindAddress := ":9090"
	l := log.New(os.Stdout, "book-store", log.LstdFlags)

	sm := http.NewServeMux()
	bh := handlers.NewBooks(l)
	sm.Handle("/book", bh)

	s := http.Server{
		Addr:         bindAddress,
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// start the server
	go func() {
		l.Println("Starting server on port ", bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	log.Println("Waiting for running tasks to complete...")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	log.Println("Shutting down")

	s.Shutdown(ctx)
}
