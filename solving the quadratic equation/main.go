package main

import (
	"fmt"
  "math"
)

func main() {
	var a, b, c float64

  fmt.Println("Введите a:")
  fmt.Scan(&a)

  fmt.Println("Введите b:")
  fmt.Scan(&b)

  fmt.Println("Введите c:")
  fmt.Scan(&c)

  d := math.Pow(b, 2) - 4 * a * c
  
  if d > 0{
    x1 := (- b + math.Sqrt(d)) / (2 * a)
    x2 := (- b - math.Sqrt(d)) / (2 * a)
    fmt.Println("Корни квадратного уравнения равны:", x1, x2)
  }
  
  if d == 0{
    x1 := - b  / (2 * a)
    fmt.Println("Корней квадратного уравнения два, но они одинаковые и равны:", x1)
  }
  
  if d < 0{
    fmt.Println("Корней нет")
  }
}
