package main

import (
	"fmt"
)

func main() {
  var number int
  fmt.Println("Введите произвольное число:")
  fmt.Scan(&number)
	
  for i := 0; i <= number; i++{
    fmt.Println(i)
  }
  fmt.Println("Достигнуто указанное число")
}
