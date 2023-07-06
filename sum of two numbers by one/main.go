package main

import (
	"fmt"
)

func main() {
  
  var number1, number2 int
  fmt.Println("Введите первое число:")
  fmt.Scan(&number1)

  fmt.Println("Введите второе число:")
  fmt.Scan(&number2)
  
  sum := number1 + number2
 
  for i := number1; i <= sum; i ++{
    fmt.Println(i)
    number1 ++
  }
  fmt.Println("Достигнута сумма введенных чисел")
}
