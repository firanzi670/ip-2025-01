package main

import (
	"fmt"
)

func main() {
	var conta int
	var consumo float64
	var tipo byte

	// Lê a entrada
	fmt.Scanf("%d %f %c", &conta, &consumo, &tipo)

	var valor float64

	switch tipo {
	case 'R': // Residencial
		valor = 5.00 + (0.05 * consumo)
	case 'C': // Comercial
		if consumo <= 80 {
			valor = 500.00
		} else {
			valor = 500.00 + 0.25*(consumo-80)
		}
	case 'I': // Industrial
		if consumo <= 100 {
			valor = 800.00
		} else {
			valor = 800.00 + 0.04*(consumo-100)
		}
	default:
		fmt.Println("Tipo de consumidor inválido.")
		return
	}

	// Imprime a saída com duas casas decimais
	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f\n", valor)
}
