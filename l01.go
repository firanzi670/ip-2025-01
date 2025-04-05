package main

import (
    "fmt"
)

func main() {
    var n1, n2, n3 float64

    // Lê as três notas separadas por espaço
    fmt.Scanf("%f %f %f", &n1, &n2, &n3)

    // Calcula a média
    media := (n1 + n2 + n3) / 3

    // Imprime a média com duas casas decimais
    fmt.Printf("MEDIA = %.2f\n", media)

    // Verifica e imprime se está aprovado ou reprovado
    if media >= 6.0 {
        fmt.Println("APROVADO")
    } else {
        fmt.Println("REPROVADO")
    }
}
