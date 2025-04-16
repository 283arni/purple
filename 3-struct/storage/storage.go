package storage

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)

	defer file.Close()

	if err != nil {
		fmt.Println("Не удалось создать файл")
	}

	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Не удалось записать файл")
	}

	fmt.Println("Успех !")
}
