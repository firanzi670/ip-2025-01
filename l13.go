package main

import (
	"fmt"
)

func main() {
	var nota float64
	fmt.Scanf("%f", &nota)

	var conceito string

	switch {
	case nota >= 9.0 && nota <= 10.0:
		conceito = "A"
	case nota >= 7.5 && nota < 9.0:
		conceito = "B"
	case nota >= 6.0 && nota < 7.5:
		conceito = "C"
	case nota >= 0.0 && nota < 6.0:
		conceito = "D"
	default:
		conceito = "Nota inválida"
	}

	// Exibe a nota com uma casa decimal e o conceito correspondente
	fmt.Printf("NOTA = %.1f CONCEITO = %s\n", nota, conceito)
}
