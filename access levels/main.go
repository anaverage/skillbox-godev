package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "test.txt"


	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()


	err = os.Chmod(fileName, 0444)
	if err != nil {
		fmt.Println("Ошибка при настройке прав доступа к файлам:", err)
		return
	}

	
	_, err = file.WriteString("Test")
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer readFile.Close()

	
	_, err = readFile.Read([]byte{})
	if err != nil {
		fmt.Println("Ошибка при чтении из файла:", err)
		return
	}

	fmt.Println("Ура, получилось!")
}
