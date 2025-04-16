package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile() {
	var name string
	fmt.Println("Введите путь до файла:")
	fmt.Scan(&name)

	_, err := os.ReadFile(name)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Файл прочтен!")
}

func CheckFile() {
	var name string
	fmt.Println("Введите путь до файла (Проверка файла на json):")
	fmt.Scan(&name)

	fileExtension := filepath.Ext(name)

	if fileExtension != ".json" {
		fmt.Println("Это не json или неверный путь")
		return
	}

	fmt.Println("Этот файл json")
}
