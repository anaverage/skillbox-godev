package main

import (
	"fmt"
	"math"
)

func main() {
	var amount, percent, year, profit float64
	fmt.Println("Введите сумму, которую хотите положить в банк:")
	fmt.Scan(&amount)
	
  fmt.Println("Введите ежемесячный процент капитализации:")
	fmt.Scan(&percent)
	
  fmt.Println("Введите количество лет, в течение которых будет открыт вклад:")
	fmt.Scan(&year)
	
  for i := 0; i < int(year) * 12; i++ {
			amount += amount * percent / 100
			_, remain := math.Modf(amount * 100)
			remain /= 100
			amount = math.Floor(amount * 100) / 100
			profit += remain
	}
	fmt.Println("Сумма, которую получит вкладчик на руки:", amount)
	fmt.Println("Выгода банка при округлении копеек равна:", profit)
}