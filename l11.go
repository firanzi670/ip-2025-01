package main

import (
    "fmt"
)

func main() {
    var numero int
    fmt.Scanln(&numero)

    if numero%3 == 0 && numero%5 == 0 {
        fmt.Println("O NUMERO E DIVISIVEL")
    } else {
        fmt.Println("O NUMERO NAO E DIVISIVEL")
    }
}
