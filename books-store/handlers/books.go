package handlers

import (
	"book-store/data"
	"log"
	"net/http"
)

// Book REST Handler
type Books struct {
	l *log.Logger
}

// Return a new Books handler
func NewBooks(l *log.Logger) *Books {
	return &Books{l}
}

// Handle books requests
func (b *Books) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	b.l.Println("handings requests for Books")

	// handle get requests
	if r.Method == http.MethodGet {
		b.getAllBooks(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		b.addBook(rw, r)
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
	return

}

// Get All Books
func (b *Books) getAllBooks(rw http.ResponseWriter, r *http.Request) {
	books := data.GetBooks()
	rw.WriteHeader(http.StatusOK)
	err := books.ToJson(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal books list to json", http.StatusInternalServerError)
	}
}

func (b *Books) addBook(rw http.ResponseWriter, r *http.Request) {
	book := &data.Book{}

	err := book.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal book from json", http.StatusBadRequest)
	}

	data.AddBook(book)
	//rw.WriteHeader(http.StatusCreated)

}
