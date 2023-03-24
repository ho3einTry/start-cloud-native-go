package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// Byte array and json in same thing in Go

func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err) // not best practice
	}
	return ToJSON
}
func FromByteArr(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var Books = map[string]Book{
	"1111111111": {Title: "Cloud Native Go", Author: "M.-L. Reamer", ISBN: "1111111111"},
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
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromByteArr(body)

		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else if len(isbn) > 5 {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method. "))

	}
}

func CreateBook(book Book) (string, bool) {

	bookExisted, exists := Books[book.ISBN]
	if exists {
		return bookExisted.ISBN, false
	}
	Books[book.ISBN] = book
	return book.ISBN, true
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
