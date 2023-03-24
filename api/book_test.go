package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJson(t *testing.T) {
	Book := Book{Title: "Cloud Native Go", Author: "M.-L. Reamer", ISBN: "0123456789"}
	json := Book.ToJSON()

	assert.Equal(t, `{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`,
		string(json), "Book Json Marshalling wrong.")
}

func TestBookFromJson(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`)
	book := FromByteArr(json)

	assert.Equal(t, Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"},
		book, "Book JSON unmarshalling wrong.")
}

func TestCreateBookIfNotExists(t *testing.T) {
	book := Book{Title: "Go learn", Author: "Parvin Zahedi", ISBN: "3333333333"}
	isbn, created := CreateBook(book)
	assert.Equal(t, true, created)
	assert.Equal(t, "3333333333", isbn)
}
func TestCreateBookIfAlreadyExists(t *testing.T) {
	book := Book{Title: "Go learn", Author: "Parvin Zahedi", ISBN: "2222222222"}
	isbn, created := CreateBook(book)
	assert.Equal(t, false, created)
	assert.Equal(t, "2222222222", isbn)
}
