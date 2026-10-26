package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/vovavdovin123-wq/simple-library/models"
)

type libraryData struct {
	Books   []*models.Book   `json:"books"`
	Readers []*models.Reader `json:"readers"`
}

func SaveLibraryToJSON(path string, lib *models.Library) error {
	data := libraryData{
		Books:   make([]*models.Book, 0, len(lib.Books)),
		Readers: make([]*models.Reader, 0, len(lib.Readers)),
	}

	for _, book := range lib.Books {
		data.Books = append(data.Books, book)
	}
	for _, reader := range lib.Readers {
		data.Readers = append(data.Readers, reader)
	}

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("не удалось закодировать JSON: %w", err)
	}

	if err := os.WriteFile(path, bytes, 0644); err != nil {
		return fmt.Errorf("не удалось записать файл: %w", err)
	}

	return nil
}

func LoadLibraryFromJSON(path string) (*models.Library, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data libraryData
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, fmt.Errorf("не удалось разобрать JSON: %w", err)
	}

	lib := models.NewLibrary()
	for _, book := range data.Books {
		lib.Books[book.ID] = book
	}
	for _, reader := range data.Readers {
		lib.Readers[reader.ID] = reader
	}

	return lib, nil
}
