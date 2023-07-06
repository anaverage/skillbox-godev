package main

import "fmt"

func main() {
    var ticket int
    fmt.Println("Введите четырёхзначный номер билета:")
    fmt.Scan(&ticket)

    ch1 := ticket / 1000
    ch2 := ticket / 100 % 10
    ch3 := ticket / 10 % 10
    ch4 := ticket % 10

    if ch1 == ch4 && ch2 == ch3 {
        fmt.Println("Зеркальный билет")
        return
    }

    if ch1 + ch2 == ch3 + ch4 {
        fmt.Println("Счастливый билет")
        return
    }

    fmt.Println("Обычный билет")
}

