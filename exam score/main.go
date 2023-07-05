package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите результат первого экзамена:")
  var firstExam int
  fmt.Scan(&firstExam)
  
  fmt.Println("Введите результат второго экзамена:")
  var secondExam int
  fmt.Scan(&secondExam)
  
  fmt.Println("Введите результат третьего экзамена:")
  var thirdExam int
  fmt.Scan(&thirdExam)
  
  fmt.Println("Сумма проходных баллов:", 275)
  
  totalScore := firstExam + secondExam + thirdExam
  
  fmt.Println("Количество набранных баллов:", totalScore)
  
  if totalScore >= 275{
    fmt.Println("Вы поступили.")
  } else {
    fmt.Println("Вы не поступили.")
  }
}
