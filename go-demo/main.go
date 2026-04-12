package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recover ", r)
		}
	}()
	fmt.Println(`___Калькулятор индекса массы тела___`)
	for {
		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			panic("Не заданы параметры для рассчета")
		}
		outputResult(IMT)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс:  %.0f", IMT)
	fmt.Println(result)
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

func calculateIMT(userKG float64, userHeight float64) (float64, error) {
	if userKG <= 0 || userHeight <= 0 {
		return 0, errors.New("Не указан вес или высота")
	}
	IMT := userKG / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
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

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Вы хотите рассчитать еще (Y/N) ? ")
	fmt.Scan(&userChoise)
	if userChoise == "Y" || userChoise == "y" {
		return true
	}
	return false
}
