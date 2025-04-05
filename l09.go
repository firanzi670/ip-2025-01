package main

import (
	"fmt"
)

func main() {
	var A, B, C float64

	// Lê os coeficientes da equação quadrática
	fmt.Scanln(&A)
	fmt.Scanln(&B)
	fmt.Scanln(&C)

	// Calcula o valor de delta
	delta := B*B - 4*A*C

	// Imprime o resultado com duas casas decimais
	fmt.Printf("O VALOR DE DELTA E = %.2f\n", delta)
}
