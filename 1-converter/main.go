package main

import "fmt"

func main() {
	const EUR float64 = 0.92
	const RUB float64 = 84.67
	sum := RUB / EUR

	fmt.Printf("%.2f", sum)
}
