package file

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileInterface interface {
	ReadFile(name string) error
	CheckFile(path string)
}

type File struct{}

func (f *File) ReadFile(name string) error {
	_, err := os.ReadFile(name)

	if err != nil {
		return err
	}

	return nil
}

func (f *File) CheckFile(path string) {
	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	var d any
	err = json.Unmarshal(data, &d)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Этот файл json")
}
