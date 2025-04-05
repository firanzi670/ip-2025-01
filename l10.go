package main

import (
	"fmt"
)

func main() {
	var a, b, c, d float64

	// Leitura dos valores da matriz 2x2
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Scanln(&d)

	// Cálculo do determinante
	determinante := a*d - b*c

	// Impressão com duas casas decimais
	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f\n", determinante)
}
