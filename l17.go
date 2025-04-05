package main

import (
    "fmt"
)

func main() {
    var x, y int
    // Lê dois inteiros da entrada
    _, err := fmt.Scanf("%d %d\n", &x, &y)
    if err != nil {
        fmt.Println("Erro na leitura da entrada.")
        return
    }

    // Verifica se o primeiro número é par
    if x%2 != 0 {
        fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
        return
    }

    // Imprime y números pares a partir de x
    for i := 0; i < y; i++ {
        fmt.Printf("%d ", x+2*i)
    }
    fmt.Println()
}
