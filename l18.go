package main

import (
	"fmt"
)

func main() {
	var a1, r, n int
	fmt.Scanf("%d %d %d", &a1, &r, &n)

	soma := 0
	for i := 0; i < n; i++ {
		termo := a1 + i*r
		soma += termo
	}

	fmt.Println(soma)
}
