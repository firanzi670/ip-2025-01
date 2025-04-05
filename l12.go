package main

import (
	"fmt"
)

func main() {
	var horas int
	fmt.Scan(&horas)

	// Calcula blocos de 3 horas
	blocosDe3h := horas / 3

	// Calcula horas restantes
	resto := horas % 3

	// Calcula valor total
	valor := float64(blocosDe3h*10 + resto*5)

	// Exibe resultado com duas casas decimais
	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valor)
}
