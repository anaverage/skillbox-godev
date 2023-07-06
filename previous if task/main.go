package main

import (
	"fmt"
)

func main() {
  var ch1, ch2, ch3 int
  
  for i := 1; i <= 1000; i ++{
    if ch1 < 10{
      ch1 ++
    }
    
    if ch2 < 100{
      ch2 ++
    }
    
    if ch3 < 1000{
      ch3 ++
    }
    fmt.Println(ch1, ch2, ch3)
  }
}  
