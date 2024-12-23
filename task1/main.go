package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/umi3uwu/Assignment1_go/library"
)

func main() {
	l := library.NewLibrary()
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter command (Add, Borrow, Return, List, Exit):")
		if !s.Scan() {
			break
		}
		c := strings.TrimSpace(s.Text())
		switch c {
		case "Add":
			fmt.Println("Enter book details in the format: ID,Title,Author")
			if !s.Scan() {
				break
			}
			i := strings.TrimSpace(s.Text())
			d := strings.Split(i, ",")
			if len(d) != 3 {
				fmt.Println("Invalid input.")
				continue
			}
			b := library.Book{
				ID:     strings.TrimSpace(d[0]),
				Title:  strings.TrimSpace(d[1]),
				Author: strings.TrimSpace(d[2]),
			}
			l.AddBook(b)
		case "Borrow":
			fmt.Println("Enter the Book ID to borrow:")
			if !s.Scan() {
				break
			}
			id := strings.TrimSpace(s.Text())
			l.BorrowBook(id)
		case "Return":
			fmt.Println("Enter the Book ID to return:")
			if !s.Scan() {
				break
			}
			id := strings.TrimSpace(s.Text())
			l.ReturnBook(id)
		case "List":
			l.ListBooks()
		case "Exit":
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid command.")
		}
	}
}
