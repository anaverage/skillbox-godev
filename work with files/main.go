package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "time"
)

func main() {
   
    file, err := os.Create("log.txt")
    if err != nil {
        fmt.Println("Не удалось открыть файл:", err)
        return
    }
    defer file.Close()

    
    writer := bufio.NewWriter(file)

   
    scanner := bufio.NewScanner(os.Stdin)

    
    count := 1

    
    for {
    
        fmt.Print("Введите сообщение (для выхода введите 'exit'): ")
        scanner.Scan()
        text := scanner.Text()

   
        if text == "exit" {
            break
        }

        t := time.Now()
        line := strconv.Itoa(count) + ". " + t.Format("2006-01-02 15:04:05") + " " + text + "\n"
        _, err := writer.WriteString(line)
        if err != nil {
            fmt.Println("Не удалось записать строку:", err)
        }

        count++
    }

    
    writer.Flush()
}
