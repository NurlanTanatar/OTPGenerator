package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i <= 10; i++ {
		fmt.Println(z)
		z = 0.5 * (z + x/z)
	}
	return z
}
func main() {
	fmt.Println("Newton's Method")
	fmt.Println(Sqrt(256))
	fmt.Println()
}
