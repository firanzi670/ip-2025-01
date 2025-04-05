package main

import (
	"fmt"
)

func fahrenheitParaCelsius(f float64) float64 {
	return (5 * (f - 32)) / 9
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var f float64
		fmt.Scan(&f)
		c := fahrenheitParaCelsius(f)
		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", f, c)
	}
}
