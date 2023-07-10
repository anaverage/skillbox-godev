package main

import (
	"fmt"
  "math"
)

func main() {
	fmt.Print("Введите степень ex: x = ")
  
  var x, n int
  fmt.Scan(&x)

  fmt.Print("Введите точность вычисления ex (количество знаков после запятой): n = ")
  fmt.Scan(&n)

  epsilon := 1 / math.Pow(10, float64(n))
  result := 1.0
  fact := 1.0
  prevResult := 0.0
  k := 1.0
  
  for {
    fact *= k
    result += math.Pow(float64(x), float64(k)) / float64(fact)
    
    if math.Abs(result - prevResult) < epsilon{
      fmt.Println("Ex равна:", result)
      break
    }
    
    k++
    prevResult = result
  }
}
