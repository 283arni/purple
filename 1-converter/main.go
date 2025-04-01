package main

import (
	"fmt"
	"strconv"
)

const USD float64 = 87.50
const EUR float64 = 92.20
const RUB float64 = 1

func scanInput() (float64, float64, float64) {
	var firstM, secondM, sum string
	var numSum float64
	var a, b float64

	for {
		fmt.Println("Введите сумму денег:")
		fmt.Scan(&sum)
		num, err := strconv.Atoi(sum)
		numSum = float64(num)

		if err != nil {
			fmt.Println("Неправельный ввод")
			continue
		}

		break
	}

	for {
		fmt.Println("Введите из какой валюты (USD, EUR, RUB):")
		fmt.Scan(&firstM)

		if firstM != "USD" && firstM != "EUR" && firstM != "RUB" {
			fmt.Println("Неправельный ввод")
			continue
		}

		if firstM == "USD" {
			a = USD
		} else if firstM == "EUR" {
			a = EUR
		} else {
			a = RUB
		}

		break
	}

	for {
		fmt.Println("Введите в какую валюту (USD, EUR, RUB):")
		fmt.Scan(&secondM)

		if secondM != "USD" && secondM != "EUR" && secondM != "RUB" {
			fmt.Println("Неправельный ввод")
			continue
		}

		if firstM == secondM {
			fmt.Println("Совпадение названий валют")
			continue
		}

		if secondM == "USD" {
			b = USD
		} else if secondM == "EUR" {
			b = EUR
		} else {
			b = RUB
		}

		break
	}

	return numSum, a, b
}

func calculate(numSum float64, a float64, b float64) float64 {
	return numSum * a / b
}

func main() {
	numSum, a, b := scanInput()
	res := calculate(numSum, a, b)

	fmt.Printf("%.2f", res)
}
