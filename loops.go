package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const Iterations = 1000
	const MinimalShift = 0.00000001

	var z float64 = 1
	for i := 0; i < Iterations; i++ {
		newVal := z - ((z*z - x) / (2 * z))
		if newVal == z || math.Abs(newVal-z) < MinimalShift {
			return z
		}
		z = newVal
	}
	return z
}

func main() {
	fmt.Println(Sqrt(123))
}
