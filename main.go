package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vovavdovin123-wq/simple-library/cmd/cli"
	"github.com/vovavdovin123-wq/simple-library/models"
	"github.com/vovavdovin123-wq/simple-library/storage"
)

const dbFile = "books.json"

func main() {
	myLibrary, err := storage.LoadLibraryFromJSON(dbFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл данных не найден, создана новая библиотека.")
			myLibrary = models.NewLibrary()
		} else {
			log.Fatalf("Ошибка при загрузке библиотеки: %v", err)
		}
	} else {
		fmt.Println("Библиотека успешно загружена из файла.")
	}

	cli.Run(myLibrary, dbFile)
}
