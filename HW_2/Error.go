package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	z := float64(1)
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(z)
		z = 0.5 * (z + x/z)
	}
	return z, nil
}
