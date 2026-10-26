package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/vovavdovin123-wq/simple-library/models"
	"github.com/vovavdovin123-wq/simple-library/storage"
)

func Run(lib *models.Library, dbPath string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("1. Список книг")
		fmt.Println("2. Список читателей")
		fmt.Println("3. Добавить книгу")
		fmt.Println("4. Добавить читателя")
		fmt.Println("5. Выдать книгу")
		fmt.Println("6. Вернуть книгу")
		fmt.Println("0. Выход")
		fmt.Print("Выберите пункт: ")

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ошибка ввода:", err)
			return
		}

		choice := strings.TrimSpace(line)
		switch choice {
		case "1":
			fmt.Println("--- Книги ---")
			for _, book := range lib.ListAllBooks() {
				fmt.Println(book)
			}

		case "2":
			fmt.Println("--- Читатели ---")
			for _, r := range lib.Readers {
				fmt.Println(r)
			}

		case "3":
			addBook(lib, reader)

		case "4":
			addReader(lib, reader)

		case "5":
			issueBook(lib, reader)

		case "6":
			returnBook(lib, reader)

		case "0":
			fmt.Println("Сохранение данных и выход.")
			if err := storage.SaveLibraryToJSON(dbPath, lib); err != nil {
				fmt.Println("Произошла ошибка сохранения списка книг:", err)
				return
			}
			fmt.Println("успех: данные сохранены в", dbPath)
			return

		default:
			fmt.Println("Неизвестный пункт меню")
		}
	}
}

func addBook(lib *models.Library, reader *bufio.Reader) {
	fmt.Print("Название: ")
	title, _ := readLine(reader)
	fmt.Print("Автор: ")
	author, _ := readLine(reader)
	fmt.Print("Год: ")
	yearStr, _ := readLine(reader)
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		fmt.Println("ошибка: неверный год")
		return
	}

	book := &models.Book{
		ID:     nextBookID(lib),
		Title:  title,
		Author: author,
		Year:   year,
	}
	if err := lib.AddBook(book); err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	fmt.Println("успех: книга добавлена")
}

func addReader(lib *models.Library, reader *bufio.Reader) {
	fmt.Print("Имя: ")
	firstName, _ := readLine(reader)
	fmt.Print("Фамилия: ")
	lastName, _ := readLine(reader)

	r := &models.Reader{
		ID:        nextReaderID(lib),
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}
	if err := lib.AddReader(r); err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	fmt.Println("успех: читатель добавлен")
}

func issueBook(lib *models.Library, reader *bufio.Reader) {
	fmt.Print("ID книги: ")
	bookID, err := readInt(reader)
	if err != nil {
		fmt.Println("ошибка: неверный ID книги")
		return
	}
	fmt.Print("ID читателя: ")
	readerID, err := readInt(reader)
	if err != nil {
		fmt.Println("ошибка: неверный ID читателя")
		return
	}

	if err := lib.IssueBookToReader(bookID, readerID); err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	fmt.Println("успех: книга выдана")
}

func returnBook(lib *models.Library, reader *bufio.Reader) {
	fmt.Print("ID книги: ")
	bookID, err := readInt(reader)
	if err != nil {
		fmt.Println("ошибка: неверный ID книги")
		return
	}

	if err := lib.ReturnBookByID(bookID); err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	fmt.Println("успех: книга возвращена")
}

func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func readInt(reader *bufio.Reader) (int, error) {
	line, err := readLine(reader)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(line)
}

func nextBookID(lib *models.Library) int {
	maxID := 0
	for id := range lib.Books {
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}

func nextReaderID(lib *models.Library) int {
	maxID := 0
	for id := range lib.Readers {
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}
