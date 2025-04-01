package main

import "fmt"

func scanInput() (int, string, string) {
	var sum int
	var firstM, secondM string
	fmt.Println("Введите сумму денег:")
	fmt.Scan(&sum)
	fmt.Println("Введите из какой валюты:")
	fmt.Scan(&firstM)
	fmt.Println("Введите в какую валюту:")
	fmt.Scan(&secondM)

	return sum, firstM, secondM
}

func calculate(sum int, a string, b string) float64 {
	return 123.4
}

func main() {
	const EUR float64 = 0.92
	const RUB float64 = 84.67

	sum, firstM, secondM := scanInput()
	res := calculate(sum, firstM, secondM)

	fmt.Printf("%.2f", res)
}
