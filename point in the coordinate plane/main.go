package main

import (
	"fmt"
)

func main() {
  var x, y int
	fmt.Println("Введите X:")
  fmt.Scan(&x)

	fmt.Println("Введите Y:")
  fmt.Scan(&y)

  if x > 0 && y > 0{
    fmt.Println("Точка принадлежит первой четверти")
  }
  if x > 0 && y < 0{
    fmt.Println("Точка принадлежит четвертой четверти")
  }
  if x < 0 && y < 0{
    fmt.Println("Точка принадлежит третьей четверти")
  
  } 
  if x < 0 && y > 0 {
    fmt.Println("Точка принадлежит второй четверти")
    }
}