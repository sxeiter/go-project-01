package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var UserHeigth float64
	var UserWeight float64
	fmt.Println("__Калькулятор индекса массы тела__")
	fmt.Println("Введите ваш рост в метрах")
	fmt.Scan(&UserHeigth)
	fmt.Println("Введите ваш вес")
	fmt.Scan(&UserWeight)
	IMT := UserWeight / math.Pow(UserHeigth, IMTPower)
	fmt.Printf("Ваш индекс массы тела: %v", math.Round(IMT))
}
