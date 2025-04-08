package main

import (
	"errors"
	"fmt"
	"strconv"
)

var money = map[string]float64{"USD": 87.50, "EUR": 92.20, "RUB": 1}

func scanInput() (float64, float64, float64) {
	var firstM, secondM, sum string
	var numSum float64
	var a, b float64

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

		val, ok := money[firstM]

		if !ok {
			fmt.Println("Введите валюту из предложеного перед вводом")
			continue
		}

		a = val

		break
	}

	for {
		fmt.Println("Введите в какую валюту (USD, EUR, RUB):")
		fmt.Scan(&secondM)

		val, ok := money[secondM]

		if !ok {
			fmt.Println("Введите валюту из предложеного перед вводом")
			continue
		}

		if firstM == secondM {
			fmt.Println("Не возможно сделать расчет так как вы ввели одинаковые названия валют: ", firstM, " и ", secondM)
			continue
		}

		b = val

		break
	}

	return numSum, a, b
}

func calculate(numSum float64, a float64, b float64) (float64, error) {

	if numSum <= 0 {
		return 0, errors.New("Сумма денег должна быть больше 0. Попробуйте заново.")
	}

	if a <= 0 || b <= 0 {
		return 0, errors.New("Валюта не может быть меньше или равняться 0. Проверьте правильность курса")
	}

	return numSum * a / b, nil
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
