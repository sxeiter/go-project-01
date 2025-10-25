package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {

	for {
		fmt.Println("__Калькулятор индекса массы тела__")
		userWeight, userHeight := getUserInput()
		IMT := calculateIMT(userWeight, userHeight)
		outputResult(IMT)
		isRepeatCalculation := checkPereatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func calculateIMT(userWeigth, userHeigth float64) float64 {
	IMT := userHeigth / math.Pow(userWeigth/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userWeight float64
	var userHeigth float64
	fmt.Println("Введите ваш рост в сантиметрах")
	fmt.Scan(&userHeigth)
	fmt.Println("Введите ваш вес")
	fmt.Scan(&userWeight)
	return userHeigth, userWeight
}

func checkPereatCalculation() bool {
	var userChoise string
	fmt.Println("Хотите сделать расчет еще раз. (Y/N) ")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
