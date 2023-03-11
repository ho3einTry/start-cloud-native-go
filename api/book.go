package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err) // not best parctice
	}
	return ToJSON
}
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var Books = []Book{
	{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"},
	{Title: "Cloud Native Net", Author: "Hossein Alizadeh", ISBN: "9876543210"},
}

func BookHandlerFunc(w http.ResponseWriter, r *http.Request) {
	books, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(books)
}
