package basic

import (
	"fmt"
	"math"
)

func Exercise() {
	squareRoot(2)
	squareRootV2(2)
	squareRootV3(2)
	fmt.Println("square root", math.Sqrt(2))
}

func squareRoot(x float64) float64 {
	z := 1.0
	var prev float64 = 0.0
	diff := 0.000000000001
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(prev-z) <= diff {
			break
		}
		prev = z
	}
	return z
}

func squareRootV2(x float64) float64 {
	z := 1.0
	var prev float64 = 0.0
	diff := 0.000000000001
	for {
		z -= (z*z - x) / (2 * z)
		if math.Abs(prev-z) <= diff {
			break
		}
		prev = z
	}
	return prev
}

func squareRootV3(x float64) float64 {
	z := x / 2
	min := 0.000000000001
	prev := 0.0
	for a := math.Inf(1); a >= min; {
		z -= (z*z - x) / (2 * z)
		a = abs(prev - z)
		if a >= min {
			prev = z
		}
	}
	return prev
}

func abs(x float64) float64 {
	if x < 0 {
		return -1 * x
	}
	return x
}
