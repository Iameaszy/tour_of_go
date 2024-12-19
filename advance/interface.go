package advance

import (
	"fmt"
	"math"
)

// A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

type Abser interface {
	Abs() float64
}

func Interface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := MethodVertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())

	var i I = &T{"hello"}
	i.M()

	fmt.Println("Take any argument", getAnyValue(a))
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t *T) M() {
	fmt.Println(t.S)
}

func (t *T) N() {
	fmt.Println(t.S)
}

type any interface{}

func getAnyValue(i any) any {
	return i
}
