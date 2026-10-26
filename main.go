package main

import "fmt"

func main() {
	fmt.Println("Проект 'Простая библиотека' запущен.")

	// Создаём пустую библиотеку
	library := NewLibrary()

	// Добавляем читателей
	reader1 := &Reader{ID: 1, FirstName: "Петр", LastName: "Петров", IsActive: true}
	reader2 := &Reader{ID: 2, FirstName: "Sergey", LastName: "Meniaylo", IsActive: true}

	if err := library.AddReader(reader1); err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	if err := library.AddReader(reader2); err != nil {
		fmt.Println("ошибка:", err)
		return
	}

	// Добавляем книги
	book1 := &Book{ID: 1, Title: "Война и мир", Author: "Лев Толстой", Year: 1869}
	book2 := &Book{ID: 2, Title: "Преступление и наказание", Author: "Ф. Достоевский", Year: 1866}

	if err := library.AddBook(book1); err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	if err := library.AddBook(book2); err != nil {
		fmt.Println("ошибка:", err)
		return
	}

	// Попытка добавить ту же книгу ещё раз (такое же название и автор)
	fmt.Println("--- Дубликат книги ---")
	duplicate := &Book{ID: 3, Title: "Война и мир", Author: "Лев Толстой", Year: 1869}
	if err := library.AddBook(duplicate); err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println("успех: книга добавлена")
	}

	fmt.Println("--- Книги в библиотеке ---")
	for _, book := range library.ListAllBooks() {
		fmt.Println(book)
	}

	// Успешная выдача книги
	fmt.Println("--- Успешная выдача книги ---")
	if err := library.IssueBookToReader(1, 1); err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println("успех: книга выдана читателю")
		fmt.Println(book1)
	}

	// Попытка выдать уже выданную книгу
	fmt.Println("--- Попытка выдать уже выданную книгу ---")
	if err := library.IssueBookToReader(1, 2); err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println("успех: книга выдана читателю")
	}

	// Попытка выдать книгу несуществующему читателю
	fmt.Println("--- Попытка выдать книгу несуществующему читателю ---")
	if err := library.IssueBookToReader(2, 99); err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println("успех: книга выдана читателю")
	}

	// Успешный возврат книги
	fmt.Println("--- Успешный возврат книги ---")
	if err := library.ReturnBookByID(1); err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println("успех: книга возвращена")
		fmt.Println(book1)
	}

	// Попытка вернуть книгу, которая уже в библиотеке
	fmt.Println("--- Попытка вернуть книгу, которая уже в библиотеке ---")
	if err := library.ReturnBookByID(1); err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println("успех: книга возвращена")
	}
}
