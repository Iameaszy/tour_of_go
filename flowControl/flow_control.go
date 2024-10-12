package flowControl

import (
	"fmt"
	"math"
)

func Run() {
	fmt.Println("Flow control")
	//loop()
	control()
}

func loop() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// Optional init and post
	// While loop
	fmt.Println("Optional init and post")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
}

func control() {
	a := 1
	if a == 1 {
		fmt.Println("This is a:", a)
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
		pow(3, 3, 8),
	)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}
