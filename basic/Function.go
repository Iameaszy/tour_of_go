package basic

import "fmt"

func Function() {
	fmt.Printf("Add x + y = %d\n", add(2, 3))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
}

// Function declaration
func add(x int, y int) int {
	return x + y
}

// Multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

// Naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
