package main

import (
	"fmt"
)

func main() {
	var fahrenheit float64
	var polegadas float64

	// Leitura dos valores
	fmt.Scanln(&fahrenheit)
	fmt.Scanln(&polegadas)

	// Conversões
	celsius := (5*fahrenheit - 160) / 9
	milimetros := polegadas * 25.4

	// Impressão com duas casas decimais
	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", celsius)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f\n", milimetros)
}
