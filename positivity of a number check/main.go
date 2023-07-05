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

  if firstNumber > 0 || secondNumber > 0 || thirdNumber > 0{
    fmt.Println("Хотя бы одно число является положительным")
  } else {
    fmt.Println("Нет ни одного положительного числа")
  }
}
