package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("__Калькулятор индекса массы тела__")
	IMT := calculateIMT(getUserInput())
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Println(result)
}

func calculateIMT(userWeigth, userHeigth float64) float64 {
	IMT := userWeigth / math.Pow(userHeigth/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeigth float64
	var userWeight float64
	fmt.Println("Введите ваш рост в сантиметрах")
	fmt.Scan(&userHeigth)
	fmt.Println("Введите ваш вес")
	fmt.Scan(&userWeight)
	return userHeigth, userWeight
}
