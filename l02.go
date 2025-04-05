package main

import (
    "fmt"
)

func main() {
    var numCasos int
    fmt.Scan(&numCasos)

    // Preços dos ingressos por categoria
    precoPopular := 1.00
    precoGeral := 5.00
    precoArquibancada := 10.00
    precoCadeiras := 20.00

    for i := 1; i <= numCasos; i++ {
        var publico int
        var pPopular, pGeral, pArquibancada, pCadeiras float64

        // Leitura dos dados
        fmt.Scan(&publico, &pPopular, &pGeral, &pArquibancada, &pCadeiras)

        // Cálculo da quantidade de pessoas por categoria
        qtdPopular := float64(publico) * (pPopular / 100)
        qtdGeral := float64(publico) * (pGeral / 100)
        qtdArquibancada := float64(publico) * (pArquibancada / 100)
        qtdCadeiras := float64(publico) * (pCadeiras / 100)

        // Cálculo da renda total
        renda := qtdPopular*precoPopular +
            qtdGeral*precoGeral +
            qtdArquibancada*precoArquibancada +
            qtdCadeiras*precoCadeiras

        // Saída formatada
        fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, renda)
    }
}
