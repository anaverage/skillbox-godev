package main

import (
	"fmt"
)

func main() {
	var firstNumber, secondNumber, thirdNumber int
  fmt.Println("Введите первое число:")
  fmt.Scan(&firstNumber)

  fmt.Println("Введите второе число:")
  fmt.Scan(&secondNumber)

  fmt.Println("Введите третье число:")
  fmt.Scan(&thirdNumber)

  if firstNumber == secondNumber || firstNumber == thirdNumber || secondNumber == thirdNumber{
    fmt.Println("Два числа или более совпадают")
  } else {
    fmt.Println("Все числа разные")
  }
}
