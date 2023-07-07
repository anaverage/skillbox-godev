package main

import (
	"fmt"
)

func main() {
  count := 0
	for i := 100000; i <= 999999; i ++{
    
    ch1 := i / 100000
    ch2 := i / 10000 % 10
    ch3 := i / 1000 % 10
    ch4 := i % 1000 / 100
    ch5 := i % 100 / 10
    ch6 := i % 10
  
    if ch1 == ch6 && ch2 == ch5 && ch3 == ch4{
      count ++
      
    }
  }
  fmt.Println("Количество зеркальных билетов среди всех шестизначных номеров от 100000 до 999999:", count)
}
