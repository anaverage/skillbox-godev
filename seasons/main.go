package main

import (
	"fmt"
)

func main() {
	fmt.Println("Времена года.")
  flag := true
  for flag {
    
    var month string
    fmt.Println("Введите месяц:")
    fmt.Scan(&month)
    
    switch month {
    
    case "декабрь", "январь", "февраль":
      fmt.Println("Время года - зима")
      flag = false
    
    case "март", "апрель", "май":
      fmt.Println("Время года - весна")
      flag = false
    
    case "июнь", "июль", "август":
      fmt.Println("Время года - лето")
      flag = false
    
    case "сентябрь", "октябрь", "ноябрь":
      fmt.Println("Время года - осень")
      flag = false
    
    default:
      fmt.Println("Некорректный ввод, попробуйте снова")    
    }
  }
}
