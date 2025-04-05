package main

import (
	"fmt"
)

func main() {
	// Constantes
	const PI = 3.14159
	const CUSTO_POR_M2 = 100.00

	// Variáveis para entrada
	var raio, altura float64

	// Leitura dos valores
	fmt.Scanln(&raio)
	fmt.Scanln(&altura)

	// Cálculo das áreas
	areaCirculo := PI * raio * raio
	areaLateral := 2 * PI * raio * altura
	areaTotal := 2*areaCirculo + areaLateral

	// Cálculo do custo
	custo := areaTotal * CUSTO_POR_M2

	// Saída formatada
	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)
}
