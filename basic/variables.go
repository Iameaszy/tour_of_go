package basic

import (
	"fmt"
	"math"
	"math/cmplx"
	"time"
)

var c, python, java bool
var j, k int = 1, 2
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func Variables() {
	fmt.Println("Hello world")

	fmt.Println("The time is now", time.Now())

	// Variables
	var i int
	fmt.Println(i, c, python, java)

	// Variables initializer
	// var can be omitted
	var d, node, deno = true, false, "no!"
	fmt.Println(j, k, d, node, deno)

	// Short variable declaration
	ab := 1
	fmt.Println(ab)

	// Go basic types
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Type conversion
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// Type inference
	v := -42.0 // change me!
	fmt.Printf("v is of type %T\n", v)

	// Constants
	const FirstName, LastName string = "", ""
	fmt.Println(FirstName, LastName)

	// Numeric Constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
