package main

import (
	"fmt"
	"struct/api"
	"struct/bins"
	"struct/file"
)

func menu() int {
	var item int

	fmt.Println("Выберите пункт меню:")
	fmt.Println("1. Добавить bin")
	fmt.Println("2. Просматреть bins")
	fmt.Println("3. Чтение файла")
	fmt.Println("4. Проверить на json")

	fmt.Scan(&item)

	return item
}

func switchMenu(item int) {
	fileHandler := &file.File{}
	binsHandler := &bins.Bins{}
	bins := binsHandler.NewBins()

	switch item {
	case 1:
		AddAccount(binsHandler, bins)
	case 2:
		ShowBins(binsHandler, bins)
	case 3:
		ReadFile(fileHandler)
	case 4:
		CheckFile(fileHandler)
	default:
		return
	}
}

func AddAccount(binsHandler bins.BinsInterface, bins *bins.Bins) {
	var private, name string
	var isPrivate bool
	fmt.Println("Будет ли bin приватный Y/n?")
	fmt.Scan(&private)
	fmt.Println("Введите имя:")
	fmt.Scan(&name)

	if private == "Y" || private == "y" {
		isPrivate = true
	} else {
		isPrivate = false
	}

	bins.AddAccount(name, isPrivate)
}

func ShowBins(binsHandler bins.BinsInterface, bins *bins.Bins) {
	binsList := binsHandler.ReadBins()

	for _, v := range binsList {
		fmt.Println(v.Name)
		fmt.Println(v.CreateAt)
	}
}

func ReadFile(fileHandler file.FileInterface) {
	var name string
	fmt.Println("Введите путь до файла:")
	fmt.Scan(&name)

	err := fileHandler.ReadFile(name)

	if err != nil {
		fmt.Println("Прочтение файла не удалось")
		return
	}

	fmt.Println("Файл упешно прочитан")
}

func CheckFile(fileHandler file.FileInterface) {
	var name string
	fmt.Println("Введите путь до файла (Проверка файла на json):")
	fmt.Scan(&name)

	fileHandler.CheckFile(name)
}

func main() {
	item := menu()
	switchMenu(item)
}
