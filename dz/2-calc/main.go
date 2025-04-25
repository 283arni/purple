package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var calcMap = map[string]func([]float64) float64{
	"AMG": addAverage,
	"SUM": add,
	"MED": addMedian,
}

func scanInput() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите название операции: \n 'AMG' - вычисление среднего значения \n 'SUM' - общая сумма \n 'MED' - найти медиану \n")
	scanner.Scan()
	command := scanner.Text()
	fmt.Println("Введите целые числа через ',' для вычисления:")
	scanner.Scan()
	numsString := scanner.Text()

	return command, numsString
}

func validater(numsString string) ([]float64, error) {
	nums := strings.Split(strings.ReplaceAll(numsString, " ", ""), ",")
	numsFloat := []float64{}
	var textError string

	for _, val := range nums {
		num, err := strconv.ParseFloat(val, 64)

		if err != nil {
			textError = val + " не является числом. Попробуйте заново."
			return numsFloat, errors.New(textError)
		}

		numsFloat = append(numsFloat, num)
	}

	return numsFloat, nil
}

func add(nums []float64) float64 {
	var sum float64

	for _, v := range nums {
		sum += v
	}

	return sum
}

func addAverage(nums []float64) float64 {
	sum := add(nums)
	average := float64(sum) / float64(len(nums))

	return average
}

func addMedian(nums []float64) float64 {
	var sum float64
	sliceLen := len(nums)
	center := sliceLen / 2

	if sliceLen%2 == 0 {
		sum = (nums[center-1] + nums[center]) / 2
	} else {
		sum = nums[center]
	}

	return sum
}

func calculate(command string, nums []float64) (float64, error) {

	funcSum, ok := calcMap[command]

	if !ok {
		return 0, errors.New("Такой команды нет, выберите из списка")
	}

	sum := funcSum(nums)

	return sum, nil
}

func main() {

	command, numsString := scanInput()
	numsFloat, err := validater(numsString)

	if err != nil {
		fmt.Print(err)
		return
	}

	sum, errCalc := calculate(command, numsFloat)

	if errCalc != nil {
		fmt.Print(errCalc)
		return
	}

	fmt.Print(sum)
}
