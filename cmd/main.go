package main

import (
	"alif-hw5/commission"
	"fmt"
	"strings"
)

func main() {
	var sender_first, sender_last string
	var receiver_first, receiver_last string
	var card_numbers string
	var amount float64
	var alif_input int

	fmt.Print("Имя отправителя: ")
	fmt.Scan(&sender_first)
	fmt.Print("Фамилия отправителя: ")
	fmt.Scan(&sender_last)
	fmt.Print("Имя получателя: ")
	fmt.Scan(&receiver_first)
	fmt.Print("Фамилия получателя: ")
	fmt.Scan(&receiver_last)
	fmt.Print("Номер карты получателя: ")
	fmt.Scan(&card_numbers)
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)
	fmt.Print("Alif карта? (1—да/0—нет): ")
	fmt.Scan(&alif_input)

	isAlif := alif_input == 1
	fee, total, err := commission.Calculate(amount, isAlif)

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	maskedCard := commission.MaskCard(card_numbers)
	tran_id := commission.GetTransaction()
	current_time := commission.GetCurrentTime()

	fmt.Println("\n========= Электронный чек ===========")
	fmt.Printf("Отправитель:%s %s\n", strings.ToUpper(sender_last), strings.ToUpper(sender_first))
	fmt.Printf("Получатель: %s %s\n", strings.ToUpper(receiver_last), strings.ToUpper(receiver_first))
	fmt.Printf("Номер транзакции: %d\n", tran_id)
	fmt.Printf("Счёт зачисления:  %s\n", maskedCard)
	fmt.Printf("Дата и время: %s\n", current_time)

	fmt.Println("-----------------------------------")
	fmt.Printf("Сумма: %.0f сум\n", amount)
	fmt.Printf("Комиссия: %.0f сум\n", fee)
	fmt.Printf("Итого: %.0f сум\n", total)

	fmt.Println("-----------------------------------")
	fmt.Println("AO ALIF TECHc* Лицензия ЦБ РУз № 000010")
	fmt.Printf("Статус:  Исполнено\n") // Синий цвет
	fmt.Println("===================================")

}
