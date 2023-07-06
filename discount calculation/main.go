package main

import (
	"fmt"
)

func main() {
	
  
  i := true
  
  for i == true{
    var price, discount int
  
    fmt.Println("Введите цену товара:")
    fmt.Scan(&price)

    fmt.Println("Введите скидку в (%):")
    fmt.Scan(&discount)
    
    if discount <= 30 && price * discount / 100 <= 2000{
      fmt.Println("Ваша скидка составляет:", price * discount / 100)
      i = false
    
    } else if discount > 30{
      fmt.Println("Введена некорректная скидка. Попробуйте снова")
    
    } else if price * discount / 100 > 2000{
      fmt.Println("Ваша скидка составляет:", 2000)
      i = false
    }
  }
}
