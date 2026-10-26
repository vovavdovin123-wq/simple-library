package main

import "fmt"

func main() {
	fmt.Println("Проект 'Простая библиотека' запущен.")

	user1 := Reader{
		ID:        1,
		FirstName: "Петр",
		LastName:  "Петров",
		IsActive:  true,
	}

	book1 := Book{
		ID:       1,
		Title:    "Война и мир",
		Author:   "Лев Толстой",
		Year:     1869,
		IsIssued: false,
	}

	library := Library{
		Books:   map[int]*Book{book1.ID: &book1},
		Readers: map[int]*Reader{user1.ID: &user1},
	}
	_ = library

	book1.IssueBook(&user1)
	fmt.Println(book1)
	user1.AssignBook(&book1)

	fmt.Println("---")

	reader2 := Reader{
		ID:        2,
		FirstName: "Sergey",
		LastName:  "Meniaylo",
		IsActive:  true,
	}
	book1.IssueBook(&reader2)

	fmt.Println("---")

	book1.ReturnBook()
	fmt.Println(book1)

	fmt.Println("---")

	user1.Deactivate()
	fmt.Println(user1)
	fmt.Println("---")
	book1.IssueBook(&user1)
}
