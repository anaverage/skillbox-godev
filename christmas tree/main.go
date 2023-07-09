package main

import "fmt"

func main() {
    fmt.Println("Вывод ёлочки.")
    fmt.Println("Введите высоту ёлочки:")
    
    var height int
    fmt.Scanln(&height)

    for i := 1; i <= height; i++ {

        for j := 1; j <= height - i; j++ {
            fmt.Print(" ")
        }

        for k := 1; k <= 2 * i - 1; k++ {
            fmt.Print("*")
        }

        fmt.Println()
    }
}
