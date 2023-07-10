package main
 
import (
    "fmt"
    "math"
)
 
func main() {
    count8 := 0
    count16 := 0
    for i := 0; i <= math.MaxUint32; i++ {
        if i > math.MaxUint8 {
            count8++
        }
        if i > math.MaxUint16 {
            count16++
        }
    }
    fmt.Println("Переполнений чисел типа uint8 в диапазоне от 0 до uint32:", count8)
    fmt.Println("Переполнений чисел типа uint16 в диапазоне от 0 до uint32:", count16)
}