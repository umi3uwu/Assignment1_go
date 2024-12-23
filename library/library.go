package library

import "fmt"

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func NewLibrary() *Library {
	return &Library{
		Books: make(map[string]Book),
	}
}

func (l *Library) AddBook(book Book) {
	if _, exists := l.Books[book.ID]; exists {
		fmt.Println("Book with this ID already exists.")
		return
	}
	l.Books[book.ID] = book
	fmt.Println("Book added successfully!")
}

func (l *Library) BorrowBook(id string) {
	book, exists := l.Books[id]
	if !exists {
		fmt.Println("Book with given ID not found.")
		return
	}
	if book.IsBorrowed {
		fmt.Println("Book is already borrowed.")
		return
	}
	book.IsBorrowed = true
	l.Books[id] = book
	fmt.Println("Book borrowed successfully!")
}

func (l *Library) ReturnBook(id string) {
	book, exists := l.Books[id]
	if !exists {
		fmt.Println("Book with given ID not found.")
		return
	}
	if !book.IsBorrowed {
		fmt.Println("Book is not currently borrowed.")
		return
	}
	book.IsBorrowed = false
	l.Books[id] = book
	fmt.Println("Book returned successfully!")
}

func (l *Library) ListBooks() {
	fmt.Println("List of available (not borrowed) books:")
	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}
