package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Negative numbers not allowed: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	const iterations = 1000
	const minimalShift = 0.00000001

	var z float64 = 1
	for i := 0; i < iterations; i++ {
		newVal := z - ((z*z - x) / (2 * z))
		if math.Abs(newVal-z) < minimalShift {
			return z, nil
		}
		z = newVal
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(5))
	fmt.Println(Sqrt(-200))
}
