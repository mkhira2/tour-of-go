package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Sqrt approximation of %v attempt %v = %v\n", x, i, z)
	}
	return z
}

func main() {
	x := 1525.0
	fmt.Printf("\n*** Newton Sqrt approximation for %v = %v\n", x, Sqrt(x))
	fmt.Printf("*** math.Sqrt response for %v = %v\n", x, math.Sqrt(x))
}
