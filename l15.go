package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	if N > 5 && N < 2000 {
		for i := 2; i <= N; i += 2 {
			fmt.Printf("%dˆ2 = %d\n", i, i*i)
		}
	}
}
