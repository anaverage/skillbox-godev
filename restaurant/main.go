package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите день недели:")
  var weekDay int
  fmt.Scan(&weekDay)

  fmt.Println("Введите число гостей:")
  var guestsCount int
  fmt.Scan(&guestsCount)

  fmt.Println("Введите сумму чека:")
  var sum int
  fmt.Scan(&sum)

  discountMonday := 0 
  
  if weekDay == 1{
    discountMonday = sum / 10
    fmt.Println("Скидка по понедельникам:", discountMonday)
    
  }

  discountFriday := 0
  
  if weekDay == 5 && sum > 10000{
    discountFriday = sum / 20
    fmt.Println("Скидка по пятницам:", discountFriday)
    
  }

  service := 0
  
  if guestsCount > 5{
    service = sum / 10
    fmt.Println("Надбавка на обслуживание:", service)
    
  }

  totalSum := sum - discountMonday - discountFriday + service
  
  fmt.Println("Сумма к оплате:", totalSum)
}
