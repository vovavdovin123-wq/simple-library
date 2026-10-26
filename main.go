package main

import "fmt"

func main() {
	fmt.Println("Проект 'Простая библиотека' запущен.")

	user1 := &Reader{
		ID:        1,
		FirstName: "Петр",
		LastName:  "Петров",
		IsActive:  true,
	}

	user2 := &Reader{
		ID:        2,
		FirstName: "Sergey",
		LastName:  "Meniaylo",
		IsActive:  true,
	}

	book1 := &Book{
		ID:       1,
		Title:    "Война и мир",
		Author:   "Лев Толстой",
		Year:     1869,
		IsIssued: false,
	}

	book2 := &Book{
		ID:       2,
		Title:    "Преступление и наказание",
		Author:   "Ф. Достоевский",
		Year:     1866,
		IsIssued: false,
	}

	book3 := &Book{
		ID:       3,
		Title:    "Мастер и Маргарита",
		Author:   "М. Булгаков",
		Year:     1967,
		IsIssued: false,
	}

	library := Library{
		Books: map[int]*Book{
			book1.ID: book1,
			book2.ID: book2,
			book3.ID: book3,
		},
		Readers: map[int]*Reader{
			user1.ID: user1,
			user2.ID: user2,
		},
	}

	fmt.Println("--- Все книги ---")
	library.ListAllBooks()

	fmt.Println("--- Выдача книги ---")
	library.IssueBookToReader(1, 1)

	fmt.Println("--- Все книги после выдачи ---")
	library.ListAllBooks()

	fmt.Println("--- Поиск книги ---")
	if book, ok := library.FindBookByID(2); ok {
		fmt.Println(book)
	}

	fmt.Println("--- Поиск читателя ---")
	if reader, ok := library.FindReaderByID(1); ok {
		fmt.Println(reader)
	}
}
