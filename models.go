package main

import "fmt"

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID *int
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

type Library struct {
	Books   map[int]*Book
	Readers map[int]*Reader
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("%s (%s, %d), статус %s", b.Title, b.Author, b.Year, status)
}

func (b *Book) IssueBook(reader *Reader) {
	if b.IsIssued {
		fmt.Printf("Книга %s уже кому-то выдана\n", b.Title)
		return
	}
	if !reader.IsActive {
		fmt.Printf("Читатель %s %s не активен и не может получить книгу.\n", reader.FirstName, reader.LastName)
		return
	}
	b.ReaderID = &reader.ID
	b.IsIssued = true
	fmt.Printf("Книга %s была выдана\n", b.Title)
}

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга %s и так в библиотеке\n", b.Title)
		return
	}
	b.ReaderID = nil
	b.IsIssued = false
	fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}

func (r *Reader) AssignBook(book *Book) {
	fmt.Printf("Читатель %s %s взял книгу %s\n", r.FirstName, r.LastName, book.String())
}

func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (r Reader) String() string {
	status := "активен"
	if !r.IsActive {
		status = "не активен"
	}
	return fmt.Sprintf("%s %s (ID %d), %s", r.FirstName, r.LastName, r.ID, status)
}

func (lib *Library) FindBookByID(id int) (*Book, bool) {
	book, ok := lib.Books[id]
	return book, ok
}

func (lib *Library) FindReaderByID(id int) (*Reader, bool) {
	reader, ok := lib.Readers[id]
	return reader, ok
}

func (lib *Library) IssueBookToReader(bookID int, readerID int) {
	book, ok := lib.FindBookByID(bookID)
	if !ok {
		fmt.Println("Книга не найдена")
		return
	}

	reader, ok := lib.FindReaderByID(readerID)
	if !ok {
		fmt.Println("Читатель не найден")
		return
	}

	book.IssueBook(reader)
}

func (lib *Library) ListAllBooks() {
	for _, book := range lib.Books {
		fmt.Println(book)
	}
}
