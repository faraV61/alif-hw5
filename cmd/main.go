package main

import (
	"alif-hw5/commission"
	"fmt"
)

func main() {
	var amount float64
	var isAlifInput int

	fmt.Print("Введите сумму: ")
	fmt.Scanln(&amount)

	fmt.Print("Alif карта? (1—да/0—нет): ")
	fmt.Scanln(&isAlifInput)
	isAlif := isAlifInput == 1
	fee, total, err := commission.Calculate(amount, isAlif)

	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Println("\n========= Чек ==========")
	fmt.Println("\nУслуга: ЛУНТИК ДУТИНОВИЧ")
	fmt.Printf("\nСумма: %.0f сум", amount)
	fmt.Printf("\nКомиссия: %.0f сум", fee)
	fmt.Printf("\nИтого: %.0f сум\n", total)
	fmt.Println("\nСтатус: Исполнено")

}
