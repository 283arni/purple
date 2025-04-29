package main

import (
	"flag"
	"fmt"
	"struct/api"
	"struct/bins"
	"struct/file"
)

func menu() int {
	var item int
	binsHandler := &bins.Bins{}

	fmt.Println("Выберите пункт меню:")
	fmt.Println("1. Добавить bin")
	fmt.Println("2. Просматреть bins")
	fmt.Println("3. Чтение файла")
	fmt.Println("4. Проверить на json")
	create := flag.Bool("create", false, "Создать бин")
	update := flag.Bool("update", false, "Обновить бин")
	file := flag.String("file", "storage/data.json", "Файл")
	id := flag.String("id", "", "ИД")
	name := flag.String("name", "", "Название")
	flag.Parse()

	if *create {
		AddAccount(binsHandler, *name, *file)
	}

	if *update {
		UpdateBin(binsHandler, *id, *file)
	}

	// fmt.Scan(&item)

	return item
}

func switchMenu(item int) {
	fileHandler := &file.File{}
	// binsHandler := &bins.Bins{}
	// bins := bins.NewBins()

	switch item {
	case 1:
		// AddAccount(binsHandler, *name, "sd" )
	case 2:
		// ShowBins(binsHandler, bins)
	case 3:
		ReadFile(fileHandler)
	case 4:
		CheckFile(fileHandler)
	default:
		return
	}
}

func AddAccount(binsHandler bins.BinsInterface, name string, file string) {

	binRes := api.PostBin(name)
	binsHandler.AddAccount(binRes, file)
}

func UpdateBin(binsHandler bins.BinsInterface, id string, file string) {
	if id == "" {
		fmt.Println("добавьте флаг --id")
	}

	bArr := binsHandler.ReadBins()

	var item bins.Bin
	for _, v := range bArr {
		if v.Metadata.Id == id {
			item = v
		}
	}

	api.UpdateBin(item)
}

func ShowBins(binsHandler bins.BinsInterface, bins *bins.Bins) {
	// binsList := binsHandler.ReadBins()

	// for _, v := range binsList {
	// 	fmt.Println(v.Name)
	// 	fmt.Println(v.CreateAt)
	// }
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
