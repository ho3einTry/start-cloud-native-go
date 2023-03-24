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

var Books = map[string]Book{
	"1111111111": {Title: "Cloud Native Go", Author: "M.-L. Reamer", ISBN: "1111111111"},
	"2222222222": {Title: "Cloud Native Net", Author: "Hossein Alizadeh", ISBN: "2222222222"},
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
func writeJson(w http.ResponseWriter, i interface{}) {
	bytes, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(bytes)
}

// func BookHandlerFunc(w http.ResponseWriter, r *http.Request) {
// 	books, err := json.Marshal(Books)
// 	if err != nil {
// 		panic(err)
// 	}
// 	w.Header().Add("Content-Type", "application/json; charset=utf-8")
// 	w.Write(books)
// }

func BooksHandlerFunc(w http.ResponseWriter, r *http.Request) {
	//isbn := r.URL.Query().Get("isbn")

	switch method := r.Method; method {
	case http.MethodGet:
		getRequest(w, r)
	case http.MethodPost:
		postRequest(w, r)
	case http.MethodPut:
		putRequest(w, r)
	case http.MethodDelete:
		deleteRequest(w, r)

	default:
		defaultRequest(w)

	}
}

func defaultRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Unsupported request method. "))
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]

	_, exists := GetBook(isbn)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
	} else {
		DeleteBook(isbn)
		w.WriteHeader(http.StatusOK)
	}

}

func putRequest(w http.ResponseWriter, r *http.Request) {

	isbn := r.URL.Path[len("/api/books/"):]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil || len(isbn) <= 1 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	book := FromByteArr(body)
	exists := UpdateBook(isbn, book)
	if exists {
		w.Header().Add("Location", "/api/books/"+isbn)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func postRequest(w http.ResponseWriter, r *http.Request) {
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
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]
	if len(isbn) >= 1 {
		book, found := GetBook(isbn)
		if found {

			writeJson(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	} else {
		books := GetAllBooks()
		writeJson(w, books)
	}
}

func DeleteBook(isbn string) {
	delete(Books, isbn)
}

func UpdateBook(isbn string, book Book) bool {
	_, exists := Books[isbn]
	if exists && isbn == book.ISBN {
		Books[isbn] = book
		return true
	}
	return false

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
