package main

import (
	"flag"
	"fmt"
	"struct/api"
	"struct/bins"
	"struct/file"
)

func AddAccount(binsHandler bins.BinsInterface, name string, file string) {

	binRes := api.PostBin(name)
	binsHandler.AddAccount(binRes, file)
}

func DeleteAccount(binsHandler bins.BinsInterface, id string) {
	if id == "" {
		panic("Обязательно укажите id")
	}

	binRes := api.DeleteBin(id)
	binsHandler.DeleteBin(binRes)
}

func ShowBins(binsHandler bins.BinsInterface) {
	binsList := binsHandler.ReadBins()

	for _, v := range binsList {
		fmt.Println(v.Id)
		fmt.Println(v.Name)
		fmt.Println(v.CreatedAt)
		fmt.Println(v.Private)
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
	binsHandler := &bins.Bins{}

	create := flag.Bool("create", false, "Создать бин")
	delete := flag.Bool("delete", false, "Обновить бин")
	list := flag.Bool("list", false, "Показать список")
	file := flag.String("file", "storage/data.json", "Файл")
	id := flag.String("id", "", "Индификатор")
	name := flag.String("name", "Имя не указано", "Название")
	flag.Parse()

	if *create {
		AddAccount(binsHandler, *name, *file)
	}

	if *delete {
		DeleteAccount(binsHandler, *id)
	}

	if *list {
		ShowBins(binsHandler)
	}
}
