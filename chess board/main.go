package main

import (
    "fmt"
)

func main() {
    var width, height int
    
    fmt.Println("Шахматная доска.")
    fmt.Println("Введите ширину:")
    fmt.Scan(&width)
    fmt.Println("Введите высоту:")
    fmt.Scan(&height)

    for row := 0; row < height; row++ {
        for col := 0; col < width; col++ {
            if (row + col) % 2 == 0 {
                fmt.Print("* ")
            } else {
                fmt.Print("  ")
            }
        }
        fmt.Println()
    }
}
