package main

import "fmt"

func main() {
	
	var char1, char2 int16
	fmt.Println("Введите первое число:")
	fmt.Scan(&char1)
	
  fmt.Println("Введите второе число:")
	fmt.Scan(&char2)
	
  var x int32
	x = int32(char1) * int32(char2)
	fmt.Println("Результат умножения первого числа на второе равен:", x)
	fmt.Print("Результат: ")
	
  if x < 0 {
		
    switch {
		case x > -127:
			fmt.Print("int8")
		case x > -32768:
			fmt.Print("int16")
		default:
			fmt.Print("int32")
		}
	
  } else {
		
    switch {
		case x < 256:
			fmt.Print("uint8")
		case x < 65536:
			fmt.Print("uint16")
		default:
			fmt.Print("uint32")
		}
	}
}