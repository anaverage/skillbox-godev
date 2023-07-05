package main

import "fmt"

func main() {
    var cost, coin1, coin2, coin3 int
    
    fmt.Print("Введите стоимость товара: ")
    fmt.Scanln(&cost)
    
    fmt.Print("Введите номинал первой монеты: ")
    fmt.Scanln(&coin1)
    
    fmt.Print("Введите номинал второй монеты: ")
    fmt.Scanln(&coin2)
    
  
    fmt.Print("Введите номинал третьей монеты: ")
    fmt.Scanln(&coin3)

    if cost == coin1 || cost == coin2 || cost == coin3 {
        fmt.Println("Можно заплатить без сдачи")
    
    } else if cost == coin1 + coin2 || cost == coin1 + coin3 || cost == coin2 + coin3 {
        fmt.Println("Можно заплатить двумя монетами")
    
    } else if cost == coin1 + coin2 + coin3 {
        fmt.Println("Можно заплатить тремя монетами")
    
    } else {
        fmt.Println("Нельзя заплатить без сдачи")
    }
}

