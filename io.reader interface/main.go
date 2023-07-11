package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		
    if os.IsNotExist(err) { 
			fmt.Println("Файл не найден")
			return
		}
		fmt.Println(err)
		return
	}
	
	if len(file) == 0 {
		fmt.Println("Файл пуст")
		return
	}
	
	fmt.Println(string(file))
}
