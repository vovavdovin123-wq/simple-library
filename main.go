package main

import "fmt"

func main() {
	fmt.Println("Проект 'Простая библиотека' запущен.")

	book := Book{
		ID:       1,
		Title:    "Война и мир",
		Author:   "Л. Толстой",
		Year:     1869,
		IsIssued: false,
	}

	fmt.Println(book)
	book.IssueBook()
	book.IssueBook()
	book.ReturnBook()
}
