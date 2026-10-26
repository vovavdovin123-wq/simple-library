package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/vovavdovin123-wq/simple-library/models"
)

func SaveBooksToCSV(path string, books []*models.Book) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Title", "Author", "Year", "IsIssued", "ReaderID"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("не удалось записать заголовок: %w", err)
	}

	for _, book := range books {
		readerID := ""
		if book.ReaderID != nil {
			readerID = strconv.Itoa(*book.ReaderID)
		}

		record := []string{
			strconv.Itoa(book.ID),
			book.Title,
			book.Author,
			strconv.Itoa(book.Year),
			strconv.FormatBool(book.IsIssued),
			readerID,
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("не удалось записать книгу: %w", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("ошибка записи CSV: %w", err)
	}

	return nil
}

func LoadBooksFromCSV(path string) ([]*models.Book, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать CSV: %w", err)
	}

	if len(records) < 2 {
		return []*models.Book{}, nil
	}

	books := make([]*models.Book, 0)
	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) < 6 {
			return nil, fmt.Errorf("строка %d: недостаточно столбцов", i+1)
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("строка %d: неверный ID: %w", i+1, err)
		}

		year, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, fmt.Errorf("строка %d: неверный Year: %w", i+1, err)
		}

		isIssued, err := strconv.ParseBool(record[4])
		if err != nil {
			return nil, fmt.Errorf("строка %d: неверный IsIssued: %w", i+1, err)
		}

		var readerID *int
		if record[5] != "" {
			value, err := strconv.Atoi(record[5])
			if err != nil {
				return nil, fmt.Errorf("строка %d: неверный ReaderID: %w", i+1, err)
			}
			readerID = &value
		}

		book := &models.Book{
			ID:       id,
			Title:    record[1],
			Author:   record[2],
			Year:     year,
			IsIssued: isIssued,
			ReaderID: readerID,
		}
		books = append(books, book)
	}

	return books, nil
}
