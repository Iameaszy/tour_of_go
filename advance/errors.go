package advance

import (
	"fmt"
	"math"
	"time"
)

func Errors() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	value, err := squareRootV3(-3)
	fmt.Println(value, err)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number %v", float64(err))
}

func squareRootV3(x float64) (float64, ErrNegativeSqrt) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

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
	return prev, 0
}

func abs(x float64) float64 {
	if x < 0 {
		return -1 * x
	}
	return x
}
