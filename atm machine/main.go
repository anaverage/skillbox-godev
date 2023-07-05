package main

import (
	"fmt"
)

func main() {
	var amount int
	fmt.Println("Введите сумму для снятия:")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Сумма должна быть положительным числом.")
	} else if amount > 100000 {
		fmt.Println("Превышен лимит суммы для снятия.")
	} else if amount%100 != 0 {
		fmt.Println("Сумма должна быть кратна 100 рублям.")
	} else {
		fmt.Println("Операция успешно выполнена.")
    fmt.Println("Вы сняли", amount, "рублей.")
	}
}
