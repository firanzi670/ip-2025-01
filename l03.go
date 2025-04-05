package main

import (
	"fmt"
	"os"
)

func main() {
	var n1, n2, n3 int

	// Leitura dos três números
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&n3)

	// Verifica se todos são dígitos (0 a 9)
	if n1 >= 0 && n1 <= 9 && n2 >= 0 && n2 <= 9 && n3 >= 0 && n3 <= 9 {
		numero := n1*100 + n2*10 + n3
		quadrado := numero * numero
		fmt.Printf("%d, %d\n", numero, quadrado)
	} else {
		fmt.Println("DIGITO INVALIDO")
		os.Exit(0)
	}
}
