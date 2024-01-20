package data

import (
	"encoding/json"
	"io"
)

type Book struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
}

type Books []*Book
 
func GetBooks() Books {
	return bookList
}

func AddBook(b *Book) {
	b.ID = getNextID()
	bookList = append(bookList, b)
}

func (b *Books) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(b)
}

func (b *Book) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(b)
}

func getNextID() int {
	b := bookList[len(bookList)-1]
	return b.ID + 1

}

var bookList = []*Book{
	{
		ID:          1,
		Name:        "The Phoenix Project",
		Author:      "The Author",
		Description: "",
		Price:       2.45,
		SKU:         "id1bk1",
	},
	{
		ID:          2,
		Name:        "The DevOps Handbook",
		Author:      "The Author",
		Description: "",
		Price:       3.45,
		SKU:         "id2bk2",
	},
}
