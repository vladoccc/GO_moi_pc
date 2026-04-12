package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("%d\n", i)
	}

	fmt.Println(`___Калькулятор индекса массы тела___`)
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
	switch {
	case IMT < 16:
		fmt.Println("У вас сильный дефицит массы")
	case IMT < 18.5:
		fmt.Println("Дефицит массы тема")
	case IMT < 25:
		fmt.Println("Номальный вес")
	case IMT < 30:
		fmt.Println("Избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс:  %.0f", IMT)
	fmt.Println(result)
}

func calculateIMT(userKG float64, userHeight float64) float64 {
	IMT := userKG / math.Pow(userHeight/100, IMTPower)
	return IMT
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите рост в см: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}
