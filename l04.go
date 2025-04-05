package main

import (
	"fmt"
)

func main() {
	var salarioMinimo, qtdKW float64

	// Leitura dos dados de entrada
	fmt.Scan(&salarioMinimo)
	fmt.Scan(&qtdKW)

	// Cálculo do custo por kW
	custoPorKW := (salarioMinimo * 0.70) / 100

	// Cálculo do custo do consumo
	custoConsumo := custoPorKW * qtdKW

	// Cálculo do custo com desconto de 10%
	custoComDesconto := custoConsumo * 0.90

	// Impressão dos resultados
	fmt.Printf("Custo por kW: R$ %.2f\n", custoPorKW)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custoConsumo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoComDesconto)
}
