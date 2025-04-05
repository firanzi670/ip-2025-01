package main

import (
	"fmt"
)

func main() {
	var horas, minutos, segundos int

	// Leitura das entradas
	fmt.Scanln(&horas)
	fmt.Scanln(&minutos)
	fmt.Scanln(&segundos)

	// Conversão para segundos
	totalSegundos := horas*3600 + minutos*60 + segundos

	// Saída formatada
	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", totalSegundos)
}
