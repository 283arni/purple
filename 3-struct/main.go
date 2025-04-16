package main

import (
	"fmt"
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
	bins := bins.NewBins()

	switch item {
	case 1:
		bins.AddAccount()
	case 2:
		bins.ReadBins()
	case 3:
		file.ReadFile()
	case 4:
		file.CheckFile()
	default:
		return
	}
}

func main() {
	item := menu()
	switchMenu(item)

}
