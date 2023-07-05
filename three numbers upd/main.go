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

  score := 0

  if number1 >= 5 {
    score ++
  }
  if number2 >= 5{
    score ++
  }
  if number3 >= 5{
    score ++
  }
  fmt.Println("Среди введённых чисел",  score, "больше или равны 5.")
}
