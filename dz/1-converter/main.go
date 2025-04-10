package main

import (
	"errors"
	"fmt"
	"strconv"
)

var money = map[string]float64{"USD": 87.50, "EUR": 92.20, "RUB": 1}

func scanInput() (float64, string, string) {
	var firstM, secondM, sum string
	var numSum float64

	for {
		fmt.Println("Введите сумму денег:")
		fmt.Scan(&sum)
		num, err := strconv.ParseFloat(sum, 64)
		numSum = float64(num)

		if err != nil {
			fmt.Println("Введите корректное число")
			continue
		}

		break

	}

	for {
		fmt.Println("Введите из какой валюты (USD, EUR, RUB):")
		fmt.Scan(&firstM)

		_, ok := money[firstM]

		if !ok {
			fmt.Println("Введите валюту из предложеного перед вводом")
			continue
		}

		break

	}

	for {
		fmt.Println("Введите в какую валюту (USD, EUR, RUB):")
		fmt.Scan(&secondM)

		_, ok := money[secondM]

		if !ok {
			fmt.Println("Введите валюту из предложеного перед вводом")
			continue
		}

		if firstM == secondM {
			fmt.Println("Не возможно сделать расчет так как вы ввели одинаковые названия валют: ", firstM, " и ", secondM)
			continue
		}

		break
	}

	return numSum, firstM, secondM
}

func calculate(numSum float64, a string, b string) (float64, error) {

	if numSum <= 0 {
		return 0, errors.New("Сумма денег должна быть больше 0. Попробуйте заново.")
	}

	if money[a] <= 0 || money[b] <= 0 {
		return 0, errors.New("Валюта не может быть меньше или равняться 0. Проверьте правильность курса")
	}

	return numSum * money[a] / money[b], nil
}

func main() {
	numSum, a, b := scanInput()
	res, err := calculate(numSum, a, b)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Результат конвертации: %.2f", res)
}
