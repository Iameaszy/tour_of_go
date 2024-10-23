package basic

import "fmt"

func Loop() {
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
