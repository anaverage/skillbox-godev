package main

import (
	"fmt"
)

func main() {
	fmt.Println("Дни недели.")

  flag := true
  for flag{
    var day string
    fmt.Println("Введите будний день недели: пн, вт, ср, чт, пт:")
    fmt.Scan(&day)

    switch day{
      
      case "пн":
        fmt.Println("понедельник")
        flag = false
        fallthrough
      
      case "вт":
        fmt.Println("вторник")
        flag = false
        fallthrough
      
      case "ср":
        fmt.Println("среда")
        fallthrough
      
      case "чт":
        fmt.Println("четверг") 
        flag = false
        fallthrough
      
      case "пт":
        fmt.Println("пятница")
        flag = false
        break
      
      default:
        fmt.Println("Некорректный ввод, попробуйте снова")
        
        
    }
  }
}
