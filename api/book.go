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

var Books = map[string]Book{
	"1111111111": {Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "1111111111"},
	"2222222222": {Title: "Cloud Native Net", Author: "Hossein Alizadeh", ISBN: "2222222222"},
}

// func BookHandlerFunc(w http.ResponseWriter, r *http.Request) {
// 	books, err := json.Marshal(Books)
// 	if err != nil {
// 		panic(err)
// 	}
// 	w.Header().Add("Content-Type", "application/json; charset=utf-8")
// 	w.Write(books)
// }

func BookHandlerFunc(w http.ResponseWriter, r *http.Request) {

}

func BooksHandlerFunc(w http.ResponseWriter, r *http.Request) {
	//isbn := r.URL.Query().Get("isbn")
	isbn := r.URL.Path[len("/api/books/"):]

	switch method := r.Method; method {
	case http.MethodGet:

		if len(isbn) >= 1 {
			book, found := GetBook(isbn)
			if found {
				//w.WriteHeader(http.StatusFound)
				writeJson(w, book)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}

		} else {
			books := GetAllBooks()
			writeJson(w, books)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method. "))

	}
}

func GetBook(isbn string) (*Book, bool) {
	book, found := Books[isbn]
	return &book, found
}

func GetAllBooks() []Book {
	arrBooks := make([]Book, len(Books))
	i := 0
	for _, book := range Books {
		arrBooks[i] = book
		i++
	}
	return arrBooks
}
func writeJson(w http.ResponseWriter, i interface{}) {
	bytes, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(bytes)
}
