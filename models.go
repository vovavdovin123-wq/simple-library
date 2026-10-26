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

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]*Book),
		Readers: make(map[int]*Reader),
	}
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("%s (%s, %d), статус %s", b.Title, b.Author, b.Year, status)
}

func (b *Book) IssueBook(reader *Reader) error {
	if reader == nil {
		return fmt.Errorf("читатель не указан")
	}
	if b.IsIssued {
		return fmt.Errorf("книга %s уже выдана", b.Title)
	}
	if !reader.IsActive {
		return fmt.Errorf("читатель %s %s не активен", reader.FirstName, reader.LastName)
	}

	b.ReaderID = &reader.ID
	b.IsIssued = true
	return nil
}

func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга %s уже в библиотеке", b.Title)
	}

	b.ReaderID = nil
	b.IsIssued = false
	return nil
}

func (r *Reader) AssignBook(book *Book) error {
	if book == nil {
		return fmt.Errorf("книга не указана")
	}
	return book.IssueBook(r)
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

func (lib *Library) AddBook(book *Book) error {
	if book == nil {
		return fmt.Errorf("книга не указана")
	}

	for _, existing := range lib.Books {
		if existing.Title == book.Title && existing.Author == book.Author {
			return fmt.Errorf("книга уже есть в библиотеке: %s (%s)", book.Title, book.Author)
		}
	}

	if _, exists := lib.Books[book.ID]; exists {
		return fmt.Errorf("книга с ID %d уже есть", book.ID)
	}

	lib.Books[book.ID] = book
	return nil
}

func (lib *Library) AddReader(reader *Reader) error {
	if reader == nil {
		return fmt.Errorf("читатель не указан")
	}
	if _, exists := lib.Readers[reader.ID]; exists {
		return fmt.Errorf("читатель с ID %d уже есть", reader.ID)
	}

	lib.Readers[reader.ID] = reader
	return nil
}

func (lib *Library) FindBookByID(id int) (*Book, error) {
	book, ok := lib.Books[id]
	if !ok {
		return nil, fmt.Errorf("книга с ID %d не найдена", id)
	}
	return book, nil
}

func (lib *Library) FindReaderByID(id int) (*Reader, error) {
	reader, ok := lib.Readers[id]
	if !ok {
		return nil, fmt.Errorf("читатель с ID %d не найден", id)
	}
	return reader, nil
}

func (lib *Library) IssueBookToReader(bookID int, readerID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}

	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}

	return book.IssueBook(reader)
}

func (lib *Library) ReturnBookByID(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	return book.ReturnBook()
}

func (lib *Library) ListAllBooks() []*Book {
	books := make([]*Book, 0, len(lib.Books))
	for _, book := range lib.Books {
		books = append(books, book)
	}
	return books
}
