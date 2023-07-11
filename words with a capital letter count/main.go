package main

import (
	"fmt"
  "unicode"
)

func main() {
	fmt.Println("Определение количества слов, начинающихся с большой буквы в строке.")
  s := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"
  count := 0
  for _, i := range s{
    if unicode.IsUpper(i){
      count ++
    }
  }
  fmt.Println("Строка содержит", count, "слов с большой буквы")
}
