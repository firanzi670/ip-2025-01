package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a float64

	// Lê os dois valores da entrada
	fmt.Scanf("%f %f", &h, &a)

	// Calcula a área da base do hexágono regular
	Ab := (3 * a * a * math.Sqrt(3)) / 2

	// Calcula o volume da pirâmide
	V := (1.0 / 3.0) * Ab * h

	// Imprime o resultado com duas casas decimais
	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", V)
}
