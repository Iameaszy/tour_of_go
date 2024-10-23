package basic

import "fmt"

func Defer() {
	basicDefer()
	deferOrder()
}

func basicDefer() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func deferOrder() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
