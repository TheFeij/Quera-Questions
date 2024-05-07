package library

import (
	"fmt"
	"strings"
)

// Book represents a book
type Book struct {
	title      string
	borrower   string
	isBorrowed bool
}

// NewBook returns a new book instance with the given title
func NewBook(title string) *Book {
	return &Book{
		title:      title,
		borrower:   "",
		isBorrowed: false,
	}
}

// Title is getter for the field title
func (b *Book) Title() string {
	return b.title
}

// Borrower is getter for the field borrower
func (b *Book) Borrower() string {
	return b.borrower
}

// IsBorrowed is getter for the field isBorrowed
func (b *Book) IsBorrowed() bool {
	return b.isBorrowed
}

// Borrow borrows the book to the given borrower
// returns false if the book is already borrowed
func (b *Book) Borrow(borrower string) bool {
	if b.isBorrowed {
		return false
	}

	b.isBorrowed = true
	b.borrower = borrower
	return true
}

// Return returns the book
// returns false if the book had not been borrowed
func (b *Book) Return() bool {
	if !b.isBorrowed {
		return false
	}

	b.isBorrowed = false
	b.borrower = ""
	return true
}

// Library represents a library that can store books up to a given capacity
type Library struct {
	len      int
	capacity int
	books    []*Book
}

// NewLibrary returns a pointer to a Library instance with the given capacity
func NewLibrary(capacity int) *Library {
	return &Library{
		books:    make([]*Book, capacity),
		len:      0,
		capacity: capacity,
	}
}

// searchBook searches for a book with the given name and returns it if exists
// returns nil if book does not exist
func (library *Library) searchBook(title string) *Book {
	title = strings.ToLower(title)
	for _, book := range library.books {
		if book != nil && book.Title() == title {
			return book
		}
	}

	return nil
}

// AddBook adds a new book to the library
func (library *Library) AddBook(name string) string {
	if book := library.searchBook(name); book != nil {
		return "The book is already in the library"
	}
	if library.len >= library.capacity {
		return "Not enough capacity"
	}

	library.books = append(library.books, NewBook(name))
	library.len += 1
	return "OK"
}

// BorrowBook peronName borrows the bookName
func (library *Library) BorrowBook(bookName, personName string) string {
	book := library.searchBook(bookName)
	if book == nil {
		return "The book is not defined in the library"
	}

	if book.IsBorrowed() {
		return "The book is already borrowed by "
	}

	if ok := book.Borrow(personName); !ok {
		return fmt.Sprintf("The book is already borrowed by %s", book.Borrower())
	}

	return "OK"
}

// ReturnBook returns a book with the given bookName to the library
func (library *Library) ReturnBook(bookName string) string {
	book := library.searchBook(bookName)
	if book == nil {
		return "The book is not defined in the library"
	}

	if !book.IsBorrowed() {
		return "The book has not been borrowed"
	}

	if ok := book.Return(); !ok {
		return "The book has not been borrowed"
	}

	return "OK"
}
