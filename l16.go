package main

import (
	"fmt"
)

func main() {
	var salario float64

	// Lê o salário da entrada padrão
	fmt.Scanln(&salario)

	var salarioReajustado float64

	// Aplica o reajuste conforme a regra
	if salario <= 300.00 {
		salarioReajustado = salario * 1.50
	} else {
		salarioReajustado = salario * 1.30
	}

	// Imprime o resultado com duas casas decimais
	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", salarioReajustado)
}
