package main

import "fmt"

func main() {
	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
		isRepeatTransaction := checkInput()
		if !isRepeatTransaction {
			break
		}
	}
	fmt.Printf("Сумма значений транзакций: %.2f", sumTransaction(transactions))
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Введите транзакцию: ")
	fmt.Scan(&transaction)
	return transaction
}

func checkInput() bool {
	var userChoise string
	fmt.Print("Вы хотите ввести значение ещё (Y/N)? ")
	fmt.Scan(&userChoise)
	if userChoise == "Y" || userChoise == "y" {
		return true
	}
	return false
}
func sumTransaction(transactions []float64) float64 {
	sum := 0.0
	for _, value := range transactions {
		sum += value
	}
	return sum
}