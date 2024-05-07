package library

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	lib := NewLibrary(100)
	res := lib.AddBook("golestan")
	assert.Equal(t, "OK", res)

	res = lib.BorrowBook("Golestan", "Ali")
	assert.Equal(t, "OK", res)

	res = lib.BorrowBook("Golestan", "Abolfazl")
	assert.Equal(t, "The book is already borrowed by Ali", res)

	res = lib.BorrowBook("Bostan", "Ali")
	assert.Equal(t, "The book is not defined in the library", res)

	res = lib.ReturnBook("golestan")
	assert.Equal(t, "OK", res)

	res = lib.ReturnBook("Bostan")
	assert.Equal(t, "The book is not defined in the library", res)

	res = lib.ReturnBook("golestan")
	assert.Equal(t, "The book has not been borrowed", res)

}
