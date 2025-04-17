package file

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(name string) error {
	_, err := os.ReadFile(name)

	if err != nil {
		return err
	}

	return nil
}

func CheckFile(path string) {
	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("Не получилось почитать файл")
		return
	}

	var d any
	err = json.Unmarshal(data, &d)

	if err != nil {
		fmt.Println("Это не json или неверный путь")
		return
	}

	fmt.Println("Этот файл json")
}
