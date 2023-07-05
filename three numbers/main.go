package main

import (
	"fmt"
)

func main() {
	fmt.Println("Три числа.")

  fmt.Println("Введите первое число:")
  var number1 int
  fmt.Scan(&number1)

  fmt.Println("Введите второе число:")
  var number2 int
  fmt.Scan(&number2)

  fmt.Println("Введите третье число:")
  var number3 int
  fmt.Scan(&number3)

  if number1 > 5 || number2 > 5 || number3 > 5{
    fmt.Println("Среди введённых чисел есть число больше 5.")
  } else {
    fmt.Println("Среди введённых чисел нет числа больше 5.")
  }
}
