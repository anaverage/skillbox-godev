package main

import (
	"fmt"
  "unicode"
  "strconv"
  
)

func main() {
	fmt.Println("Исходная строка:")
  s := "a10 10 20b 20 30c30 30 dd"
  fmt.Println(s)
  
  var nums []string
  var curNum string
    
  for _, i := range s {
    if unicode.IsDigit(i) {
      curNum += string(i)
    } else {
        if curNum != "" {
          nums = append(nums, curNum)
          curNum = ""
        }
      }
  }
    
  if curNum != "" {
        nums = append(nums, curNum)
  }
    
  fmt.Print("В строке содержатся числа в десятичном формате: ")
    
  for _, numStr := range nums {
    if num, err := strconv.Atoi(numStr); err == nil {
      fmt.Printf("%d ", num)
    }
  }
    
  fmt.Println()
}
