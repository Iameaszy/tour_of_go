package generics

import "fmt"

func Main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(findIndex(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(findIndex(ss, "hello"))

	l1 := &List[int]{
		next: &List[int]{
			next: &List[int]{
				next: nil, val: 5,
			},
			val: 4,
		},
		val: 3,
	}
	for {
		if l1 == nil {
			break
		}
		fmt.Println("linked list", l1.val)
		l1 = l1.next
	}
}

// Index returns the index of x in s, or -1 if not found.
func findIndex[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}
