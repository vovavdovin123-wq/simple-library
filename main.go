package main

import (
	"fmt"
	"os"

	"github.com/vovavdovin123-wq/simple-library/models"
	"github.com/vovavdovin123-wq/simple-library/storage"
)

func main() {
	library := models.NewLibrary()

	if err := seedLibrary(library); err != nil {
		fmt.Println("ошибка:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Использование:")
		fmt.Println("  go run . list")
		fmt.Println("  go run . save <имя_файла.csv>")
		fmt.Println("  go run . load <имя_файла.csv>")
		return
	}

	command := os.Args[1]

	switch command {
	case "list":
		fmt.Println("--- Книги в библиотеке ---")
		for _, book := range library.ListAllBooks() {
			fmt.Println(book)
		}

	case "save":
		if len(os.Args) != 3 {
			fmt.Println("Использование: go run . save <имя_файла.csv>")
			return
		}
		path := os.Args[2]
		if err := storage.SaveBooksToCSV(path, library.ListAllBooks()); err != nil {
			fmt.Println("ошибка:", err)
			return
		}
		fmt.Println("успех: книги сохранены в", path)

	case "load":
		if len(os.Args) != 3 {
			fmt.Println("Использование: go run . load <имя_файла.csv>")
			return
		}
		path := os.Args[2]
		books, err := storage.LoadBooksFromCSV(path)
		if err != nil {
			fmt.Println("ошибка:", err)
			return
		}
		library.ReplaceBooks(books)
		fmt.Println("успех: книги загружены из", path)
		fmt.Println("--- Книги после загрузки ---")
		for _, book := range library.ListAllBooks() {
			fmt.Println(book)
		}

	default:
		fmt.Println("Неизвестная команда:", command)
	}
}

func seedLibrary(library *models.Library) error {
	readers := []*models.Reader{
		{ID: 1, FirstName: "Петр", LastName: "Петров", IsActive: true},
		{ID: 2, FirstName: "Sergey", LastName: "Meniaylo", IsActive: true},
	}
	for _, reader := range readers {
		if err := library.AddReader(reader); err != nil {
			return err
		}
	}

	books := []*models.Book{
		{ID: 1, Title: "Война и мир", Author: "Лев Толстой", Year: 1869},
		{ID: 2, Title: "Преступление и наказание", Author: "Ф. Достоевский", Year: 1866},
		{ID: 3, Title: "Мастер и Маргарита", Author: "М. Булгаков", Year: 1967},
	}
	for _, book := range books {
		if err := library.AddBook(book); err != nil {
			return err
		}
	}

	return nil
}
