package advance

import (
	"fmt"
	"strings"
)

func Array() {
	// arrayDeclaration()
	//slices()
	//slicesWithMake()
	//slicesOfSlices()
	loopingByRange()
}

func arrayDeclaration() {
	// Method 1
	var a [2]string
	a[0] = "Hello world"
	a[1] = "Hi"

	for i := 0; i < len(a); i++ {
		fmt.Println("array", a[i])
	}

	// Method 2
	vowels := [5]string{"a", "e", "i", "o", "u"}
	fmt.Println(vowels)

	// method 3
	// Compiler count array lenght
	nums := [...]int{1, 2, 3, 4}
	printSlice(nums[:])
}

func slices() {
	vowels := [5]string{"a", "e", "i", "o", "u"}
	// s is a slice
	var s []string = vowels[0:3]
	fmt.Println("s", s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// Array
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// Slice
	r := []bool{true, false, true, true, false, true}
	fmt.Println("fmt r, length, capacity", r, len(r), cap(r))

	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(st)

	sb := []int{2, 3, 5, 7, 11, 13}
	sc := []int{2, 3, 5, 7, 11, 13}

	sb = sb[1:4]
	fmt.Println(s)

	sb = sb[:2]
	fmt.Println(s)

	sb = sb[1:]
	fmt.Println(sb)

	sb = sb[:]
	fmt.Println(sb)

	fmt.Println(sc[:])
	//sc length
	fmt.Println("sc len", len(sc))

	// sc capacity
	fmt.Println("sc cap", cap(sc))
	sc = sc[:0]
	fmt.Println("sc cap and length", cap(sc), len(sc), sc)
	sc = sc[:4]
	fmt.Println("sc cap and length", cap(sc), len(sc), sc)

	var nilSlice []int
	if nilSlice == nil {
		fmt.Println("nilSlice is nil!")
	}
}

func slicesWithMake() {
	// Can't create nil slice with make
	integerSlice := make([]int, 0)
	fmt.Println(integerSlice, len(integerSlice), cap(integerSlice), integerSlice == nil)

	var integerSliceV2 []int
	fmt.Println(integerSliceV2, len(integerSliceV2), cap(integerSliceV2), integerSliceV2 == nil)

	integerSliceV2 = append(integerSliceV2, 0)
	printSlice(integerSliceV2)

	integerSliceV2 = append(integerSliceV2, 1, 2, 3, 4, 5)
	printSlice(integerSliceV2)
}

func slicesOfSlices() {
	board := [][]string{
		{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func loopingByRange() {
	mySlice := [][]string{
		{"2"},
		{"3"},
		{"1"},
	}

	for i, j := range mySlice {
		fmt.Println(i, j)
	}
}

func PicExercise() {

}
